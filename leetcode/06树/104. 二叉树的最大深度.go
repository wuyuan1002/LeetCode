package main

// 104. 二叉树的最大深度

// 给定一个二叉树 root ，返回其最大深度。
//
// 二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。

// maxDepth .
// leetcode 111. 二叉树的最小深度
// leetcode 102. 二叉树的层序遍历
// leetcode 543. 二叉树的直径
// leetcode 110. 平衡二叉树
//
// 后序遍历二叉树，递归求出左右子树的最大深度，然后加上当前节点的深度返回
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 分别求出左右子树的最大深度
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// 返回当前节点的最大深度+1 -- 当前节点的深度也需要被加上
	return max(leftDepth, rightDepth) + 1
}
