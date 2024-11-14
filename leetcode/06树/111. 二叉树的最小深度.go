package main

// 111. 二叉树的最小深度

// 给定一个二叉树，找出其最小深度。
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
// 说明：叶子节点是指没有子节点的节点。

// minDepth .
// leetcode 104. 二叉树的最大深度
// leetcode 102. 二叉树的层序遍历
//
// 后序遍历二叉树，递归求出左右子树的最小深度，然后加上当前节点的深度返回
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 若左子树为空则当前节点的最小深度为其右子树的最小深度+1
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	// 若右子树为空则当前节点的最小深度为其左子树的最小深度+1
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	// 若左右子树都不为空则当前节点的最小深度为其左右子树的最小深度最小值+1
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
