package main

// 226. 翻转二叉树

// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

// invertTree .
// Offer 27. 二叉树的镜像
//
// 前中后序遍历二叉树，把每个节点的左右子树交换
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left

	return root
}
