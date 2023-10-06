package main

// 104. 二叉树的最大深度

// 给定一个二叉树 root ，返回其最大深度。
//
// 二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。

// maxDepth .
// 同 leetcode 111. 二叉树的最小深度、leetcode 102. 二叉树的层序遍历
// 1. 使用102的层级遍历二叉树，遍历过程中统计二叉树一共有多少层，层数即为二叉树的最大深度
// 2. 后序遍历二叉树，递归求出左右子树的最大深度，然后加上当前节点的深度返回
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 分别求出左右子树的最大深度
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// 返回最大深度+1 -- 当前节点的深度也需要被加上
	return max(leftDepth, rightDepth) + 1
}

// max .
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
