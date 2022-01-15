package main

import (
	"container/list"
)

// 剑指 Offer 55 - I. 二叉树的深度

// 输入一棵二叉树的根节点，求该树的深度。
// 从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

// func main() {

// }

// 求出左右子树深度，比较大小返回大的那个加1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return max(leftDepth, rightDepth) + 1
}

// 第二次做 -- 本题中，一棵树的深度等于它左子树或者右子树的深度+1
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth1(root.Left)
	rightDepth := maxDepth1(root.Right)

	return max(leftDepth, rightDepth) + 1
}

// 第二次做 -- 使用广度优先遍历树，每遍历完一行，深度+1
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0 // 深度

	queue := list.List{}
	queue.PushBack(root)

	num := 1     // 当前行还有几个
	nextNum := 0 // 下一行的个数

	var node *TreeNode
	for queue.Len() > 0 {
		node = queue.Remove(queue.Front()).(*TreeNode)
		num--

		if node.Left != nil {
			queue.PushBack(node.Left)
			nextNum++
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
			nextNum++
		}

		if num == 0 {
			depth++
			num = nextNum
			nextNum = 0
		}
	}
	return depth
}
