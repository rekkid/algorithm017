//Given an array of strings strs, group the anagrams together. You can return th
//e answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters of a differe
//nt word or phrase, typically using all the original letters exactly once.
//
//
// Example 1:
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
//Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2:
// Input: strs = [""]
//Output: [[""]]
// Example 3:
// Input: strs = ["a"]
//Output: [["a"]]
//
//
// Constraints:
//
//
// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] consists of lower-case English letters.
//
// Related Topics Hash Table String
// 👍 4341 👎 207



import java.util.*;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        if (strs.length == 0) {
            return new ArrayList<>();
        }
        Map<String, List<String>> map = new HashMap<>();
        for (String s: strs) {
            char[] chars = s.toCharArray();
            Arrays.sort(chars);
            String cs = String.valueOf(chars);
            if (!map.containsKey(cs)) {
                map.put(cs, new ArrayList<String>());
            }
            map.get(cs).add(s);
        }

        return new ArrayList<>(map.values());
    }
}
//leetcode submit region end(Prohibit modification and deletion)
