//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105
//
// Related Topics 数组 双指针
// 👍 3272 👎 0

package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	var threeSumTests = []struct {
		in       []int
		expected [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}},
		{[]int{}, [][]int{}},
		{[]int{0}, [][]int{}},
		{[]int{0, 0, 0, 0}, [][]int{[]int{0, 0, 0}}},
		{[]int{3, 0, -2, -1, 1, 2}, [][]int{[]int{-2, -1, 3}, []int{-2, 0, 2}, []int{-1, 0, 1}}},
	}
	for _, tt := range threeSumTests {
		actual := threeSum(tt.in)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("threeSum(%v) = %v; expected %v", tt.in, actual, tt.expected)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)

//func threeSum3(nums []int) [][]int {
//	nums_map := make(map[int]int)
//	for i := 0; i < len(nums);i++ {
//		if _, ok := nums_map[nums[i]]; !ok {
//			nums_map[nums[i]] = i
//		}
//	}
//
//	for i := 0; i < len(nums) - 1; i++ {
//		for j := 1; j < len(nums); j++ {
//			tmp := nums[i] + nums[j]
//
//		}
//	}
//}

func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for first := 0; first < len(nums)-2; first++ {
		if nums[first] > 0 {
			break
		}
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second, third := first+1, len(nums)-1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if sum == 0 {
				result = append(result, []int{nums[first], nums[second], nums[third]})
				second++
				third--
				for second < third && nums[second] == nums[second-1] {
					second++
				}
				for second < third && nums[third] == nums[third+1] {
					third--
				}
			} else if sum < 0 {
				second++
			} else {
				third--
			}
		}
	}

	return result
}

func threeSum_my(nums []int) [][]int {
	//1. 对数组进行排序
	//2. 从两边向中间找满足条件的数
	if len(nums) <= 2 {
		return [][]int{}
	}
	sort.Ints(nums)
	fmt.Println(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				// 得到一组解
				tmp := []int{nums[i], nums[j], nums[k]}
				result = append(result, tmp)
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return result
}

// 暴力计算——超时
func threeSum_brutal(nums []int) [][]int {
	// a + b = -c
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					tmp := []int{nums[i], nums[j], nums[k]}
					if findRepeated(result, tmp) {
						continue
					}
					result = append(result, tmp)
				}
			}
		}
	}
	return result
	//return [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}
	//return [][]int{}
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func findRepeated(x [][]int, y []int) bool {
	for _, item := range x {
		sort.Ints(item)
		sort.Ints(y)
		if equal(item, y) {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
