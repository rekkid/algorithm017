//Merge two sorted linked lists and return it as a new sorted list. The new list
// should be made by splicing together the nodes of the first two lists.
//
// Example:
//
//
//Input: 1->2->4, 1->3->4
//Output: 1->1->2->3->4->4
//
// Related Topics Linked List
// 👍 4993 👎 631

//21. 合并两个有序链表
//        将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
//        示例：
//
//        输入：1->2->4, 1->3->4
//        输出：1->1->2->3->4->4
//        通过次数379,816提交次数589,562

package leetcode.editor.en;

public class MergeTwoSortedLists {
    public static void main(String[] args) {

        Solution solution = new MergeTwoSortedLists().new Solution();
//        ListNode l1 = new MergeTwoSortedLists().new ListNode(4);
//        l1 = l1.add(2);
//        l1 = l1.add(1);
//        l1.print();

        System.out.println("===========");
//        ListNode l2 = new MergeTwoSortedLists().new ListNode(4);
//        l2 = l2.add(3);
//        l2 = l2.add(1);
//        l2.print();

        System.out.println("after merge: ===========");
//        ListNode l = solution.mergeTwoLists(l1, l2);
//        l.print();
    }
    //leetcode submit region begin(Prohibit modification and deletion)

    /**
     * Definition for singly-linked list.
     * public class ListNode {
     * int val;
     * ListNode next;
     * ListNode() {}
     * ListNode(int val) { this.val = val; }
     * ListNode(int val, ListNode next) { this.val = val; this.next = next; }
     * }
     */

    public class ListNode {
        int val;
        ListNode next;

        ListNode() {
        }

        ListNode(int val) {
            this.val = val;
            this.next = null;
        }

        ListNode(int val, ListNode next) {
            this.val = val;
            this.next = next;
        }

        @Override
        public String toString() {
            return "ListNode{" +
                    "val=" + val +
                    ", next=" + next +
                    '}';
        }
////        ListNode add(int val) {
////            ListNode l = new ListNode(val);
////            l.next = this;
////            return l;
////        }
////
////        void print() {
////            ListNode l = this;
////            while (l != null) {
////                System.out.println(l.val);
////                l = l.next;
////            }
////        }
    }

    class Solution {
        public ListNode mergeTwoListsRecrusive(ListNode l1, ListNode l2) {
            if (l1 == null) {
                return l2;
            }
            if (l2 == null) {
                return l1;
            }
            if (l1.val < l2.val) {
                l1.next = mergeTwoLists(l1.next, l2);
                return l1;
            } else {
                l2.next = mergeTwoLists(l1, l2.next);
                return l2;
            }
        }

        public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
            ListNode prehead = new ListNode(-1);

            ListNode prev = prehead;
            while (l1 != null && l2 != null) {
                if (l1.val <= l2.val) {
                    prev.next = l1;
                    l1 = l1.next;
                } else {
                    prev.next = l2;
                    l2 = l2.next;
                }
                prev = prev.next;
            }
            prev.next = l1 == null ? l2 : l1;

            return prehead.next;
        }

    }
//leetcode submit region end(Prohibit modification and deletion)

}
