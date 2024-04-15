package main

// 938. 二叉搜索树的范围和

// 给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。

// rangeSumBST .
// 中序遍历二叉搜索树，当节点进入目标范围后累加结果值，出目标范围后直接返回
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	} else if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	} else if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	} else {
		return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	}
}
