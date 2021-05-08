//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//
//
// 示例 2：
//
//
//输入：nums = [0]
//输出：[[],[0]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// nums 中的所有元素 互不相同
//
// Related Topics 位运算 数组 回溯算法
// 👍 1159 👎 0

package main

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	var subsetsTests = []struct {
		in       []int
		expected [][]int
	}{
		{[]int{1, 2, 3, 4, 5}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		//{[]int{0}, [][]int{{}, {0}}},
	}
	for _, tt := range subsetsTests {
		actual := subsets(tt.in)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("subsets(%v) = %v; expected %v", tt.in, actual, tt.expected)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)

func subsets_ok(nums []int) [][]int {
	result := [][]int{}
	var gen func(n int, tmp []int)
	gen = func(n int, tmp []int) {
		if n == len(nums) {
			b := make([]int, len(tmp))
			copy(b, tmp)
			result = append(result, b)
			//result = append(result, append([]int{}, tmp...))
			return
		}

		gen(n+1, tmp)
		tmp = append(tmp, nums[n])
		gen(n+1, tmp)
		//tmp = tmp[:len(tmp) - 1]
		//return
	}
	gen(0, []int{})
	return result
}

func subsets(nums []int) [][]int {
	res := [][]int{{}}
	for _, num := range nums {
		for _, v := range res {
			tmp := append(v, num)
			res = append(res, append([]int{}, tmp...))
		}

		// !!上下两种方法等价，使用上面的即可

		//for _, v := range res {
		//	tmp := make([]int, len(v))
		//	copy(tmp, v)
		//	tmp = append(tmp, num)
		//	res = append(res, tmp)
		//}
	}
	return res
}

func subsets3(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		dfs(cur + 1)
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]

	}
	dfs(0)
	return
}

//
//func subsets2(nums []int) [][]int {
//	result =[][]int{}
//	gen(len(nums), nums, []int{})
//	return result
//}

/*func gen(n int, nums []int, tmp []int) {
	if n == 0 {
		sort.Ints(tmp)
		//b := make([]int, len(tmp))
		//copy(b, tmp)
		result = append(result, tmp)
		return
	}
	//var new []int
	//b := make([]int, len(tmp))
	//copy(b, tmp)

	tmp = append(tmp, nums[n - 1])
	//b := make([]int, len(tmp))
	//copy(b, tmp)
	gen(n - 1, nums, tmp)
	tmp = tmp[:len(tmp) - 1]
	gen(n - 1, nums, tmp)
	return
}


func subsetsMyMod(nums []int) [][]int {
	result =[][]int{}
	gen(len(nums), nums, []int{})
	return result
}

func gen1(n int, nums []int, tmp []int) {
	if n == 0 {
		sort.Ints(tmp)
		//b := make([]int, len(tmp))
		//copy(b, tmp)
		result = append(result, tmp)
		return
	}
	//var new []int
	//b := make([]int, len(tmp))
	//copy(b, tmp)
	gen(n - 1, nums, tmp)
	tmp = append(tmp, nums[n - 1])
	b := make([]int, len(tmp))
	copy(b, tmp)
	gen(n - 1, nums, b)
	return
}
*/
//leetcode submit region end(Prohibit modification and deletion)
