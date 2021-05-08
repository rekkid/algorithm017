//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
// Related Topics 字符串 回溯算法
// 👍 1763 👎 0

package main

import (
	"reflect"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	var generateParenthesisTests = []struct {
		in       int
		expected []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{1, []string{"()"}},
	}
	for _, tt := range generateParenthesisTests {
		actual := generateParenthesis(tt.in)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("generateParenthesis(%v) = %v; expected %v", tt.in, actual, tt.expected)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis_my(n int) []string {
	var result []string
	var generate func(tmp string, left int, right int, level int)
	generate = func(tmp string, left int, right int, level int) {
		if level == n*2 {
			result = append(result, tmp)
			return
		}

		if left > 0 {
			generate(tmp+"(", left-1, right, level+1)
		}
		if right > left {
			generate(tmp+")", left, right-1, level+1)
		}
		return
	}
	generate("", n, n, 0)
	return result
}

func generateParenthesis(n int) []string {
	result := generate("", 0, 0, n, []string{})
	return result
}

func generate(tmp string, left int, right int, max int, result []string) []string {
	if left == max && right == max {
		result = append(result, tmp)
		return result
	}

	if left < max {
		result = generate(tmp+"(", left+1, right, max, result)
	}
	if right < left {
		result = generate(tmp+")", left, right+1, max, result)
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
