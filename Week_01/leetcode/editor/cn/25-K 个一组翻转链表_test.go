//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 进阶：
//
//
// 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
//
//
// 示例 3：
//
//
//输入：head = [1,2,3,4,5], k = 1
//输出：[1,2,3,4,5]
//
//
// 示例 4：
//
//
//输入：head = [1], k = 1
//输出：[1]
//
//
//
//
//
// 提示：
//
//
// 列表中节点的数量在范围 sz 内
// 1 <= sz <= 5000
// 0 <= Node.val <= 1000
// 1 <= k <= sz
//
// Related Topics 链表
// 👍 1072 👎 0

package main

import (
	"testing"
)

func TestReverseNodesInKGroup(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, head}
	prev := dummy
	next := head
	for {
		count := 0
		for count < k && next != nil {
			next = next.Next
			count++
		}
		if count < k {
			return dummy.Next
		}

		prev.Next = reverseListHelper(head, next)
		head.Next = next
		prev = head
		head = next
	}
}

func reverseListHelper(head *ListNode, tail *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != tail {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reverseKGroup_recrusive(head *ListNode, k int) *ListNode {
	len := 0
	tmp := head
	for tmp != nil {
		len++
		tmp = tmp.Next
	}
	if len <= 1 || len < k {
		return head
	}

	// 可以进行至少一次反转

	var prev *ListNode
	curr := head
	count := 0
	for count < k {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		count++
	}
	head.Next = reverseKGroup(curr, k)
	return prev

}

//leetcode submit region end(Prohibit modification and deletion)
