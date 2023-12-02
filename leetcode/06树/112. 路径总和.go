package main

// 112. 路径总和

// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。
// 判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
// 如果存在，返回 true ；否则，返回 false 。
//
// 叶子节点 是指没有子节点的节点。

// hasPathSum .
// 同 leetcode 113. 路径总和 II、leetcode、437. 路径总和 III
// 前序遍历二叉树，同时相加路径上的值判断是否满足路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfsHasPathSum(root, 0, targetSum)
}

// dfsHasPathSum 前序遍历二叉树，判断给定节点是否为结果路径上的节点
func dfsHasPathSum(node *TreeNode, currentSum, targetSum int) bool {
	if node == nil {
		return false
	}

	// 将当前节点的值加到当前和上
	currentSum += node.Val

	// 若当前节点为叶子节点并且当前和为目标值，说明当前节点为结果路径上的节点
	if node.Left == nil && node.Right == nil && currentSum == targetSum {
		return true
	}

	// 若当前节点还不满足条件，则递归遍历其左右子节点
	return dfsHasPathSum(node.Left, currentSum, targetSum) || dfsHasPathSum(node.Right, currentSum, targetSum)
}
