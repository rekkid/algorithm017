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
		{[]int{2, 1, 5, 6, 2, 3}, 10},
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
type Bar struct {
	height int
	index  int
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	n := len(heights) + 2
	// Add a sentry at the beginning and the end
	getHeight := func(i int) int {
		if i == 0 || n-1 == i {
			return 0
		}
		return heights[i-1]
	}
	stack := make([]int, 0, n/2)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && getHeight(stack[len(stack)-1]) > getHeight(i) {
			// pop stack
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			maxArea = max(maxArea, getHeight(idx)*(i-stack[len(stack)-1]-1))
		}
		// push stack
		stack = append(stack, i)
	}
	return maxArea
}

func largestRectangleArea2(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func largestRectangleArea1(heights []int) int {
	if len(heights) < 1 {
		return math.MinInt64
	}
	if len(heights) == 1 {
		return heights[0]
	}
	var maxStack []Bar
	var maxArea int
	maxStack = append(maxStack, Bar{-1, -1})
	for i := 0; i < len(heights); i++ {
		currentHeight := heights[i]
		top := maxStack[len(maxStack)-1].height
		if currentHeight >= top {
			maxStack = append(maxStack, Bar{currentHeight, i})
		} else {
			// 找到第一个不大于currentHeight的位置
			for pos := len(maxStack) - 1; pos >= 0; pos-- {
				if maxStack[pos].height > currentHeight {
					area := (i - maxStack[pos-1].index - 1) * maxStack[pos].height
					if maxArea < area {
						maxArea = area
					}
					maxStack = maxStack[:len(maxStack)-1]
				} else {
					break
				}
			}
			maxStack = append(maxStack, Bar{currentHeight, i})
		}
	}

	max := maxStack[len(maxStack)-1].index
	for i := 1; i < len(maxStack); i++ {
		area := (max - maxStack[i-1].index) * maxStack[i].height
		if maxArea < area {
			maxArea = area
		}
	}

	return maxArea
}

//leetcode submit region end(Prohibit modification and deletion)
