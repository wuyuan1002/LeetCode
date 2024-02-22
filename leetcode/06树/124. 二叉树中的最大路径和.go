package main

import "math"

// 124. 二叉树中的最大路径和

// 二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
//
// 路径和 是路径中各节点值的总和。
//
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。

// maxPathSum .
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	pathSum(root, &maxSum)
	return maxSum
}

// pathSum .
// 返回给定根节点能够为其父节点提供的最大贡献值（若贡献值为负则返回0）
// 在计算的过程中更新全局最大路径和
func pathSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	// 左右子树能为该节点提供的最大贡献 -- 最小为0
	leftSum := pathSum(root.Left, maxSum)
	rightSum := pathSum(root.Right, maxSum)

	// 计算当前节点作为根节点的最大路径和
	priceNewPath := root.Val + leftSum + rightSum

	// 更新全局最大路径和
	*maxSum = max(*maxSum, priceNewPath)

	// 返回当前节点对其父节点的最大贡献值（若贡献值为负则返回0）
	return max(root.Val+max(leftSum, rightSum), 0)
}
