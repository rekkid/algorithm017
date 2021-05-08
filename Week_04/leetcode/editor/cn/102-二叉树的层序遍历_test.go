//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例：
//二叉树：[3,9,20,null,null,15,7],
//
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其层序遍历结果：
//
//
//[
//  [3],
//  [9,20],
//  [15,7]
//]
//
// Related Topics 树 广度优先搜索
// 👍 858 👎 0

package main

import "testing"

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// BFS
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	var queue []*TreeNode

	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		var tmp []int
		var nodes []*TreeNode
		for _, node := range queue {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		res = append(res, tmp)
		queue = nodes
	}
	return res
}

func levelOrder_DFS_OK(root *TreeNode) [][]int {
	res := make([][]int, 0)
	var DFS func(level int, node *TreeNode)
	DFS = func(level int, node *TreeNode) {
		if node == nil {
			return
		}
		// 判断如果当前层是新层次，则新创建一层空间
		if level == len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], node.Val)
		DFS(level+1, node.Left)
		DFS(level+1, node.Right)
	}
	DFS(0, root)
	return res
}

func levelOrder_DFS(root *TreeNode) [][]int {
	nodeMap := make(map[int][]int)
	var DFS func(level int, node *TreeNode)
	DFS = func(level int, node *TreeNode) {
		if node == nil {
			return
		}
		nodeMap[level] = append(nodeMap[level], node.Val)
		DFS(level+1, node.Left)
		DFS(level+1, node.Right)
	}
	DFS(0, root)
	res := make([][]int, len(nodeMap))
	for level, val := range nodeMap {
		res[level] = val
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
