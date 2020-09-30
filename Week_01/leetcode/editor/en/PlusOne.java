//Given a non-empty array of digits representing a non-negative integer, increme
//nt one to the integer.
//
// The digits are stored such that the most significant digit is at the head of
//the list, and each element in the array contains a single digit.
//
// You may assume the integer does not contain any leading zero, except the numb
//er 0 itself.
//
//
// Example 1:
//
//
//Input: digits = [1,2,3]
//Output: [1,2,4]
//Explanation: The array represents the integer 123.
//
//
// Example 2:
//
//
//Input: digits = [4,3,2,1]
//Output: [4,3,2,2]
//Explanation: The array represents the integer 4321.
//
//
// Example 3:
//
//
//Input: digits = [0]
//Output: [1]
//
//
//
// Constraints:
//
//
// 1 <= digits.length <= 100
// 0 <= digits[i] <= 9
//
// Related Topics Array
// 👍 1771 👎 2615
//
//66. 加一
//        给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//
//        最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
//        你可以假设除了整数 0 之外，这个整数不会以零开头。
//
//        示例 1:
//
//        输入: [1,2,3]
//        输出: [1,2,4]
//        解释: 输入数组表示数字 123。
//        示例 2:
//
//        输入: [4,3,2,1]
//        输出: [4,3,2,2]
//        解释: 输入数组表示数字 4321。
//        通过次数207,552提交次数456,566
//
package leetcode.editor.en;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

public class PlusOne{
  public static void main(String[] args) {
       Solution solution = new PlusOne().new Solution();
       int[] digits = {1, 2, 3};
       assertArrayEquals(new int[]{1, 2, 4}, solution.plusOne(digits));
  }
  //leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public int[] plusOne(int[] digits) {
        int n = digits.length;
        for (int i = n - 1; i >= 0; i--) {
            if (digits[i] < 9) {
                digits[i]++;
                return digits;
            } else {
                digits[i] = 0;
            }
        }
        int[] nums = new int[n + 1];
        nums[0] = 1;
        return nums;
    }
}
//leetcode submit region end(Prohibit modification and deletion)

}
