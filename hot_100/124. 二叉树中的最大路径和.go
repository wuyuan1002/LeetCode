package main

import (
	"math"
)

// 124. 二叉树中的最大路径和

// 路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。
// 同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
// 路径和 是路径中各节点值的总和。
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。

// func main() {

// }

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	pathSum(root, &maxSum)
	return maxSum
}

func pathSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	// 左右树能提供的最大值 -- 最小为0
	leftSum := pathSum(root.Left, maxSum)
	rightSum := pathSum(root.Right, maxSum)

	// 当前节点的最大值
	priceNewPath := root.Val + leftSum + rightSum

	// 更新最大值
	*maxSum = max(*maxSum, priceNewPath)

	// 返回当前节点对父节点的最大贡献值 -- 若贡献值小于0，则返回0
	return max(root.Val+max(leftSum, rightSum), 0)
}
