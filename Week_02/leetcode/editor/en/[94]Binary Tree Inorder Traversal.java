//Given the root of a binary tree, return the inorder traversal of its nodes' va
//lues.
//
//
// Example 1:
//
//
//Input: root = [1,null,2,3]
//Output: [1,3,2]
//
//
// Example 2:
//
//
//Input: root = []
//Output: []
//
//
// Example 3:
//
//
//Input: root = [1]
//Output: [1]
//
//
// Example 4:
//
//
//Input: root = [1,2]
//Output: [2,1]
//
//
// Example 5:
//
//
//Input: root = [1,null,2]
//Output: [1,2]
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
//
//
//
//
// Follow up:
//
// Recursive solution is trivial, could you do it iteratively?
//
//
// Related Topics Hash Table Stack Tree
// 👍 3834 👎 166


//leetcode submit region begin(Prohibit modification and deletion)


import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Deque;
import java.util.List;

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    List<Integer> list = new ArrayList<>();
    public List<Integer> inorderTraversal1(TreeNode root) {
        if (root != null) {
            list = inorderTraversal1(root.left);
            list.add(root.val);
            list = inorderTraversal1(root.right);
        }
        return list;
    }

    public List<Integer> inorderTraversal(TreeNode root) {
        if (root == null) {
            return new ArrayList<>();
        }
        Deque<TreeNode> stack = new ArrayDeque<>();
        if (root.right != null) {
            stack.push(root.right);
        }
        stack.push(root);
        if (root.left != null) {
            stack.push(root.left);
        }
        TreeNode n = stack.pop();
//        while (!stack.isEmpty() && n.right != null) {
//            if (n.right != null) {
//                stack.push(n.right);
//            }
//            stack.push(n);
//            if (n.left != null) {
//                stack.push(n.left);
//            }
//            n = stack.pop();
//        }
//
//        while (!stack.isEmpty()) {
//            n = stack.pop();
//            list.add(n.val);
//        }
        return list;
    }


}
//leetcode submit region end(Prohibit modification and deletion)
