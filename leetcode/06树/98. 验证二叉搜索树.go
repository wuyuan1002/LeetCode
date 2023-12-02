package main

import "math"

// 98. 验证二叉搜索树

// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征：
//
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。

// isValidBST .
// 1. 前序遍历，通过最大最小值限定，校验节点是否符合要求
func isValidBST(root *TreeNode) bool {
	return dfsIsValidBST(root, math.MinInt64, math.MaxInt64)
}

// dfsIsValidBST 检查给定节点的树是否满足二叉搜索树
// min: 当前节点值的最小值
// max: 当前节点值的最大值
func dfsIsValidBST(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	// 若当前节点的值不在(min, max)范围内，说明二叉树不满足二叉搜索树
	if node.Val <= min || node.Val >= max {
		return false
	}

	// 检查当前节点的左右子树是否满足二叉搜索树
	return dfsIsValidBST(node.Left, min, node.Val) && dfsIsValidBST(node.Right, node.Val, max)
}

// ----------------

// isValidBST1 .
// 同 leetcode 530. 二叉搜索树的最小绝对差
// 2. 若为二叉搜索树则中序遍历为递增的，因此，在中序遍历时校验当前值是否比前一个值大即可
func isValidBST1(root *TreeNode) bool {
	preVal := math.MinInt64
	return dfsIsValidBST1(root, &preVal)
}

// dfsIsValidBST1 检查给定的节点是否满足二叉搜索树
// preVal: 前一个访问的节点值
func dfsIsValidBST1(node *TreeNode, preVal *int) bool {
	if node == nil {
		return true
	}

	// 检查左子树是否满足二叉搜索树
	if !dfsIsValidBST1(node.Left, preVal) {
		return false
	}

	// 检查当前节点是否满足二叉搜索树
	if node.Val <= *preVal {
		return false
	}
	*preVal = node.Val

	// 检查右子树是否满足二叉搜索树
	return dfsIsValidBST1(node.Right, preVal)
}
