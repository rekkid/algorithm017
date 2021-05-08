//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 注意：给定 n 是一个正整数。
//
// 示例 1：
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//
// 示例 2：
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
//
// Related Topics 动态规划
// 👍 1636 👎 0

package main

import (
	"reflect"
	"testing"
)

func TestClimbingStairs(t *testing.T) {
	var climbStairsTests = []struct {
		in       int
		expected int
	}{
		{2, 2},
		{3, 3},
	}
	for _, tt := range climbStairsTests {
		actual := climbStairs(tt.in)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("climbStairs(%v) = %v; expected %v", tt.in, actual, tt.expected)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	first := 1
	second := 2
	for i := 2; i < n; i++ {
		tmp := second
		second = first + second
		first = tmp
	}
	return second
}

func climbStairs_withMemo(n int) int {
	memo := make(map[int]int, n+1)
	res := climbStairsHelper(n, memo)
	return res
}

func climbStairsHelper(n int, memo map[int]int) int {
	if n == 1 || n == 2 {
		return n
	}
	if _, ok := memo[n]; !ok {
		memo[n] = climbStairsHelper(n-1, memo) + climbStairsHelper(n-2, memo)
	}
	return memo[n]
}

func climbStairs_recrusive(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	res := climbStairs(n-1) + climbStairs(n-2)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
