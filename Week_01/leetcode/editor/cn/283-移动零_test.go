//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例:
//
// 输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
// 说明:
//
//
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
//
// Related Topics 数组 双指针
// 👍 1027 👎 0

package main

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	var moveZerosTests = []struct {
		in       []int // input
		expected []int // expected result
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
	}

	for _, tt := range moveZerosTests {
		actual := make([]int, len(tt.in))
		copy(actual, tt.in)
		moveZeroes(actual)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("moveZeroes(%v) = %v; expected %v", tt.in, actual, tt.expected)
		}
	}

}

//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes_brutal(nums []int) {
	count := 0
	len := len(nums)
	for i := 0; i < len-1 && count < len-1; count++ {
		if nums[i] != 0 {
			i++
			continue
		}
		for j := i + 1; j < len; j++ {
			nums[j-1] = nums[j]
		}
		nums[len-1] = 0
	}
}

func moveZeroes(nums []int) {
	first := 0
	len := len(nums)
	for second := 0; second < len; second++ {
		if nums[second] != 0 {
			nums[first], nums[second] = nums[second], nums[first]
			first++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
