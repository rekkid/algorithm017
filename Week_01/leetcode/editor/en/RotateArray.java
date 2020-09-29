//Given an array, rotate the array to the right by k steps, where k is non-negat
//ive.
//
// Follow up:
//
//
// Try to come up as many solutions as you can, there are at least 3 different w
//ays to solve this problem.
// Could you do it in-place with O(1) extra space?
//
//
//
// Example 1:
//
//
//Input: nums = [1,2,3,4,5,6,7], k = 3
//Output: [5,6,7,1,2,3,4]
//Explanation:
//rotate 1 steps to the right: [7,1,2,3,4,5,6]
//rotate 2 steps to the right: [6,7,1,2,3,4,5]
//rotate 3 steps to the right: [5,6,7,1,2,3,4]
//
//
// Example 2:
//
//
//Input: nums = [-1,-100,3,99], k = 2
//Output: [3,99,-1,-100]
//Explanation:
//rotate 1 steps to the right: [99,-1,-100,3]
//rotate 2 steps to the right: [3,99,-1,-100]
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 2 * 10^4
// It's guaranteed that nums[i] fits in a 32 bit-signed integer.
// k >= 0
//
// Related Topics Array
// 👍 3312 👎 859

/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

        示例 1:

        输入: [1,2,3,4,5,6,7] 和 k = 3
        输出: [5,6,7,1,2,3,4]
        解释:
        向右旋转 1 步: [7,1,2,3,4,5,6]
        向右旋转 2 步: [6,7,1,2,3,4,5]
        向右旋转 3 步: [5,6,7,1,2,3,4]
        示例 2:

        输入: [-1,-100,3,99] 和 k = 2
        输出: [3,99,-1,-100]
        解释:
        向右旋转 1 步: [99,-1,-100,3]
        向右旋转 2 步: [3,99,-1,-100]
        说明:

        尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
        要求使用空间复杂度为 O(1) 的 原地 算法。

        来源：力扣（LeetCode）
        链接：https://leetcode-cn.com/problems/rotate-array
        著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

package leetcode.editor.en;

public class RotateArray {
    public static void main(String[] args) {
        Solution solution = new RotateArray().new Solution();
        int[] nums = {1,2,3,4,5,6};
        solution.rotate(nums, 2);
        for (int i: nums) {
            System.out.println(i);
        }
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public void rotateBruce(int[] nums, int k) {
            k = k % nums.length;
            for (int i = 0; i < k; i++) {
                int temp = nums[nums.length - 1];
                for (int j = nums.length - 1; j > 0; j--) {
                    nums[j] = nums[j - 1];
                }
                nums[0] = temp;
            }
        }

        public void rotateBruce2(int[] nums, int k) {
            k = k % nums.length;
            int temp = 0;
            int previous = 0;
            for (int i = 0; i < k; i++) {
                previous = nums[nums.length - 1];
                for (int j = 0; j < nums.length; j++) {
                    temp = nums[j];
                    nums[j] = previous;
                    previous = temp;
                }
            }
        }

        public void rotateCycle(int[] nums, int k) {
            k = k % nums.length;
            int count = 0;
            for (int start = 0; count < nums.length; start++) {
                int current = start;
                int prev = nums[start];
                do {
                    int next = (current + k) % nums.length;
                    int temp = nums[next];
                    nums[next] = prev;
                    prev = temp;
                    current = next;
                    count++;
                } while (start != current);
            }
        }

        public void reverse(int[] nums, int start, int end) {
            int i = start;
            int j = end;
            while (i < j) {
                int temp = nums[i];
                nums[i] = nums[j];
                nums[j] = temp;
                i++;
                j--;
            }
        }

        public void rotate(int[] nums, int k) {
            k = k % nums.length;
            reverse(nums, 0, nums.length - 1);
            reverse(nums, 0, k - 1);
            reverse(nums, k, nums.length - 1);
        }
    }
//leetcode submit region end(Prohibit modification and deletion)

}
