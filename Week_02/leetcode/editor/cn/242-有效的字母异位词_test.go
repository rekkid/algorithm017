//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 示例 1:
//
// 输入: s = "anagram", t = "nagaram"
//输出: true
//
//
// 示例 2:
//
// 输入: s = "rat", t = "car"
//输出: false
//
// 说明:
//你可以假设字符串只包含小写字母。
//
// 进阶:
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
// Related Topics 排序 哈希表
// 👍 377 👎 0

package main

import (
	"reflect"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	var isAnagramTests = []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagara4", true},
		//{"rat", "car", false},
	}
	for _, tt := range isAnagramTests {
		actual := isAnagram(tt.s, tt.t)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("isAnagram(%v, %v) = %v; expected %v", tt.s, tt.t, actual, tt.expected)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charMap := make(map[rune]int)
	for _, char := range s {
		charMap[char]++
	}
	for _, char := range t {
		// !!下面3行多余
		if _, ok := charMap[char]; !ok {
			return false
		}
		charMap[char]--
		if charMap[char] < 0 {
			return false
		}
	}
	return true
}

func isAnagram_my(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := make(map[rune]int)
	for _, char := range s {
		if _, ok := charMap[char]; !ok {
			// !!此处易错
			charMap[char] = 1
		} else {
			charMap[char]++
		}
	}

	for _, char := range t {
		if _, ok := charMap[char]; !ok {
			return false
		} else {
			charMap[char]--
		}
	}

	for _, v := range charMap {
		if v != 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
