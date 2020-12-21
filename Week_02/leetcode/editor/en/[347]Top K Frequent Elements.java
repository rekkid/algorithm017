//Given a non-empty array of integers, return the k most frequent elements. 
//
// Example 1: 
//
// 
//Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
// 
//
// 
// Example 2: 
//
// 
//Input: nums = [1], k = 1
//Output: [1] 
// 
//
// Note: 
//
// 
// You may assume k is always valid, 1 ≤ k ≤ number of unique elements. 
// Your algorithm's time complexity must be better than O(n log n), where n is t
//he array's size. 
// It's guaranteed that the answer is unique, in other words the set of the top 
//k frequent elements is unique. 
// You can return the answer in any order. 
// 
// Related Topics Hash Table Heap 
// 👍 4134 👎 243


import java.util.Comparator;
import java.util.HashMap;
import java.util.Map;
import java.util.PriorityQueue;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        // 统计每个数值出现次数
        Map<Integer, Integer> map = new HashMap<>();
        for (int num: nums) {
            map.put(num, map.getOrDefault(num, 0) + 1);
        }

        // 使用堆获得前k高频率的数值

        PriorityQueue<int[]> queue = new PriorityQueue<>(new Comparator<int[]>() {
            @Override
            public int compare(int[] m, int[] n) {
                return m[1] - n[1];
            }
        });
        for (Map.Entry<Integer, Integer> entry: map.entrySet()) {
            int num = entry.getKey();
            int occurrence = entry.getValue();
            if (queue.size() < k) {
                queue.offer(new int[]{num, occurrence});
            } else {
                int[] item = queue.peek();
                if (item[1] < occurrence) {
                    queue.poll();
                    queue.offer(new int[]{num, occurrence});
                }
            }
        }

        // 将堆中的结果保存到数值中并返回
        int[] result = new int[k];
        int i = 0;
        for (int[] item: queue) {
            result[i] = item[0];
            i++;
        }
        return result;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
