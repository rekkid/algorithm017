//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
//
//
// 示例 1：
//
//
//输入：s = "()"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "()[]{}"
//输出：true
//
//
// 示例 3：
//
//
//输入：s = "(]"
//输出：false
//
//
// 示例 4：
//
//
//输入：s = "([)]"
//输出：false
//
//
// 示例 5：
//
//
//输入：s = "{[]}"
//输出：true
//
//
//
// 提示：
//
//
// 1 <= s.length <= 104
// s 仅由括号 '()[]{}' 组成
//
// Related Topics 栈 字符串
// 👍 2377 👎 0

package main

import (
	"reflect"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	var isValidTests = []struct {
		in       string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"){", false},
	}
	for _, tt := range isValidTests {
		actual := isValid(tt.in)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("isValid(%v) = %v; expected %v", tt.in, actual, tt.expected)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	stack := []rune{}
	pair := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
	}
	if len(s)&1 != 0 {
		return false
	}
	for _, ch := range s {
		if ch == '[' || ch == '(' || ch == '{' {
			stack = append(stack, ch)
		} else {
			if _, ok := pair[ch]; !ok {
				return false
			}
			// !!陷阱一定先判断stack长度不为0
			if len(stack) > 0 && stack[len(stack)-1] == pair[ch] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
