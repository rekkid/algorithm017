//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
// 你可以按任意顺序返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//
//
// 示例 2：
//
//
//输入：nums = [3,2,4], target = 6
//输出：[1,2]
//
//
// 示例 3：
//
//
//输入：nums = [3,3], target = 6
//输出：[0,1]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 103
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// 只会存在一个有效答案
//
// Related Topics 数组 哈希表
// 👍 10971 👎 0

package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}
	var threeSumTests = []struct {
		input
		expected []int
	}{
		{input{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{input{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{input{[]int{3, 3}, 6}, []int{0, 1}},
	}
	for _, tt := range threeSumTests {
		actual := twoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("threeSum(%v) = %v; expected %v", tt.input, actual, tt.expected)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		pos, ok := numsMap[target-nums[i]]
		if !ok {
			numsMap[nums[i]] = i
		} else {
			return []int{pos, i}
		}
	}
	return []int{}
}

func twoSum_map_two_times(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		pos, ok := numsMap[target-nums[i]]
		// 要注意判断不能与自己位置相等
		if !ok || pos == i {
			continue
		}
		return []int{i, pos}
	}
	return []int{}
}

func twoSum_brutal(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//leetcode submit region end(Prohibit modification and deletion)
