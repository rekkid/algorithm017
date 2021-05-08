//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串
// 👍 732 👎 0

package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	var groupAnagramsTests = []struct {
		in       []string
		expected [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{
			[]string{"ate", "eat", "tea"},
			[]string{"nat", "tan"},
			[]string{"bat"},
		}},
	}
	for _, tt := range groupAnagramsTests {
		actual := groupAnagrams(tt.in)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("groupAnagrams(%v) = %v; expected %v", tt.in, actual, tt.expected)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	wordMap := make(map[string][]string)
	for _, str := range strs {
		word := []byte(str)
		//fmt.Println(string(word))
		sort.Slice(word, func(i, j int) bool {
			return word[i] < word[j]
		})
		//fmt.Println("after sort")
		//fmt.Println(string(word))
		wordMap[string(word)] = append(wordMap[string(word)], str)
	}

	result := [][]string{}
	for _, v := range wordMap {
		result = append(result, v)
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)
