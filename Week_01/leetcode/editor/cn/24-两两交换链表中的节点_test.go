//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//
// 示例 2：
//
//
//输入：head = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
//
//
//
//
// 进阶：你能在不修改链表节点值的情况下解决这个问题吗?（也就是说，仅修改节点本身。）
// Related Topics 递归 链表
// 👍 904 👎 0

package main

import "testing"

func TestSwapNodesInPairs(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs_iterative(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	tmp := dummyHead
	for tmp.Next != nil && tmp.Next.Next != nil {
		first := tmp.Next
		second := tmp.Next.Next

		first.Next = second.Next
		second.Next = first
		tmp.Next = second

		tmp = first
	}
	return dummyHead.Next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

func swapPairs_my(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := new(ListNode)
	dummyHead.Next = head

	tmp := dummyHead
	first := head
	second := head.Next

	head = second

	for second != nil {
		first.Next = second.Next
		second.Next = first
		tmp.Next = second

		if first.Next == nil {
			break
		}
		tmp = first
		first = first.Next
		second = first.Next
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
