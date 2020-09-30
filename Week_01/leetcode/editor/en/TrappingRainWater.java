//Given n non-negative integers representing an elevation map where the width of
// each bar is 1, compute how much water it is able to trap after raining.
//
//
//The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In
//this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos
// for contributing this image!
//
// Example:
//
//
//Input: [0,1,0,2,1,0,1,3,2,1,2,1]
//Output: 6
// Related Topics Array Two Pointers Stack
// 👍 8282 👎 130


//42. 接雨水
//        给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
//        上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。
//
//        示例:
//
//        输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//        输出: 6
//        通过次数152,729提交次数289,938

package leetcode.editor.en;

import java.util.Deque;
import java.util.LinkedList;

public class TrappingRainWater {
    public static void main(String[] args) {
        Solution solution = new TrappingRainWater().new Solution();
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {

        // 使用单调栈
        public int trapStack(int[] height) {
            int ans = 0, current = 0;
            Deque<Integer> stack = new LinkedList<Integer>();
            while (current < height.length) {
                while (!stack.isEmpty() && height[current] > height[stack.peek()]) {
                    int top = stack.pop();
                    if (stack.isEmpty())
                        break;
                    int distance = current - stack.peek() - 1;
                    int bounded_height = Math.min(height[current], height[stack.peek()]) - height[top];
                    ans += distance * bounded_height;
                }
                stack.push(current++);
            }
            return ans;
        }

        // 使用双指针
        public int trap(int[] height) {
            int left = 0, right = height.length - 1;
            int ans = 0;
            int left_max = 0, right_max = 0;
            while (left < right) {
                if (height[left] < height[right]) {
                    if (height[left] >= left_max) {
                        left_max = height[left];
                    } else {
                        ans += (left_max - height[left]);
                    }
                    ++left;
                } else {
                    if (height[right] >= right_max) {
                        right_max = height[right];
                    } else {
                        ans += (right_max - height[right]);
                    }
                    --right;
                }
            }
            return ans;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)

}
