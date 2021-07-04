package main

import (
	"math"
)

// 124. 二叉树中的最大路径和

// 路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。
// 同一个节点在一条路径序列中至多出现一次 。该路径至少包含一个节点，且不一定经过根节点。
// 路径和是路径中各节点值的总和。
//
// 给你一个二叉树的根节点root，返回其最大路径和。

func main() {

}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := math.MinInt32
	pathSum1(root, &maxSum)
	return maxSum
}

func pathSum1(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	leftSum := pathSum1(node.Left, maxSum)
	rightSum := pathSum1(node.Right, maxSum)

	sum := node.Val + leftSum + rightSum

	*maxSum = max(*maxSum, sum)

	return max(node.Val+max(leftSum, rightSum), 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
