package main

import "math"

// 530. 二叉搜索树的最小绝对差

// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
//
// 差值是一个正数，其数值等于两值之差的绝对值。

// getMinimumDifference .
// 同 leetcode 98. 验证二叉搜索树
// 中序遍历二叉树，过程中记录最小差值
// 二叉搜索树的中序遍历是递增的，相当于是求一个递增序列的最小差值，递增序列的最小差值一定是两个相邻数的差值
func getMinimumDifference(root *TreeNode) int {
	preVal, result := math.MaxInt64, math.MaxInt64
	dfsGetMinimumDifference(root, &preVal, &result)
	return result
}

// dfsGetMinimumDifference 中序遍历二叉树，求最小差值
// preVal: 遍历的前一个节点
// result: 最小差值
func dfsGetMinimumDifference(node *TreeNode, preVal *int, result *int) {
	if node == nil {
		return
	}

	// 遍历左子树
	dfsGetMinimumDifference(node.Left, preVal, result)

	// 使用当前节点值和前一个节点的值求最小差值
	if *preVal != math.MaxInt64 {
		*result = min(*result, node.Val-*preVal)
	}

	// 标记当前节点为已遍历
	*preVal = node.Val

	// 遍历右子树
	dfsGetMinimumDifference(node.Right, preVal, result)
}
