
//Given a sorted array nums, remove the duplicates in-place such that each eleme
//nt appears only once and returns the new length.
//
// Do not allocate extra space for another array, you must do this by modifying
//the input array in-place with O(1) extra memory.
//
// Example 1:
//
//
//Given nums = [1,1,2],
//
//Your function should return length = 2, with the first two elements of nums be
//ing 1 and 2 respectively.
//
//It doesn't matter what you leave beyond the returned length.
//
// Example 2:
//
//
//Given nums = [0,0,1,1,1,2,2,3,3,4],
//
//Your function should return length = 5, with the first five elements of nums b
//eing modified to 0, 1, 2, 3, and 4 respectively.
//
//It doesn't matter what values are set beyond the returned length.
//
//
// Clarification:
//
// Confused why the returned value is an integer but your answer is an array?
//
// Note that the input array is passed in by reference, which means a modificati
//on to the input array will be known to the caller as well.
//
// Internally you can think of this:
//
//
//// nums is passed in by reference. (i.e., without making a copy)
//int len = removeDuplicates(nums);
//
//// any modification to nums in your function would be known by the caller.
//// using the length returned by your function, it prints the first len element
//s.
//for (int i = 0; i < len; i++) {
//    print(nums[i]);
//}
// Related Topics Array Two Pointers
// 👍 2925 👎 5739
//
//26. 删除排序数组中的重复项
//        给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//
//        不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//
//
//
//        示例 1:
//
//        给定数组 nums = [1,1,2],
//
//        函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
//
//        你不需要考虑数组中超出新长度后面的元素。
//        示例 2:
//
//        给定 nums = [0,0,1,1,1,2,2,3,3,4],
//
//        函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
//
//        你不需要考虑数组中超出新长度后面的元素。
//
//
//        说明:
//
//        为什么返回数值是整数，但输出的答案是数组呢?
//
//        请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//
//        你可以想象内部操作如下:
//
//// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
//        int len = removeDuplicates(nums);
//
//// 在函数里修改输入数组对于调用者是可见的。
//// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
//        for (int i = 0; i < len; i++) {
//        print(nums[i]);
//        }
//        通过次数434,141提交次数837,977
//

package leetcode.editor.en;

public class RemoveDuplicatesFromSortedArray {
    public static void main(String[] args) {
        Solution solution = new RemoveDuplicatesFromSortedArray().new Solution();
        int[] nums = {0, 0, 1, 1, 1, 2, 2, 3, 3, 4};
        int len = solution.removeDuplicates(nums);
        for (int i = 0; i < len; i++) {
            System.out.println(nums[i]);
        }

    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int removeDuplicates(int[] nums) {
            if (nums.length == 0)
                return 0;
            int i = 0;
            for (int j = 1; j < nums.length; j++) {
                if (nums[j] != nums[i]) {
                    i++;
                    nums[i] = nums[j];
                }
            }

            return i + 1;
        }
    }
//leetcode submit region end(Prohibit modification and deletion)

}
