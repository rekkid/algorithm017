//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
//
//
// 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
//
//
//
//
//
// 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
//
//
//
// 示例:
//
// 输入: [2,1,5,6,2,3]
//输出: 10
// Related Topics 栈 数组
// 👍 1347 👎 0

package main

import (
	"math"
	"reflect"
	"testing"
)

func TestLargestRectangleInHistogram(t *testing.T) {
	var largestRectangleAreaTests = []struct {
		height   []int
		expected int
	}{
		//{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 1, 2}, 3},
	}
	for _, tt := range largestRectangleAreaTests {
		actual := largestRectangleArea(tt.height)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("largestRectangleArea(%v) = %v; expected %v", tt.height, actual, tt.expected)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
type Area struct {
	height int
	width  int
}

func largestRectangleArea(heights []int) int {
	if len(heights) < 1 {
		return math.MinInt64
	}
	if len(heights) == 1 {
		return heights[0]
	}
	var maxStack []Area
	var maxArea int
	maxStack = append(maxStack, Area{heights[0], 1})
	for i := 1; i < len(heights); i++ {
		currentHeight := heights[i]
		top := maxStack[len(maxStack)-1].height
		if currentHeight >= top {
			maxStack = append(maxStack, Area{currentHeight, 0})
			for j, _ := range maxStack {
				maxStack[j].width++
			}
		} else {
			for j := len(maxStack) - 1; j >= 0; j-- {
				if maxStack[j].height > currentHeight {
					area := maxStack[j].width * maxStack[j].height
					if maxArea < area {
						maxArea = area
					}

					maxStack = maxStack[:len(maxStack)-1]

					maxStack = append(maxStack, Area{currentHeight, 1})
				} else {
					maxStack[j].width++
					maxStack = append(maxStack, Area{currentHeight, 1})
					break
				}
			}
			maxStack = append(maxStack, Area{currentHeight, 1})
		}
	}

	for _, v := range maxStack {
		area := v.width * v.height
		if maxArea < area {
			maxArea = area
		}
	}

	return maxArea
}

//leetcode submit region end(Prohibit modification and deletion)
