package main

// 701. 二叉搜索树中的插入操作

// 给定二叉搜索树（BST）的根节点 root 和要插入树中的值 value ，将值插入二叉搜索树。
// 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。
//
// 注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。你可以返回 任意有效的结果 。

// insertIntoBST .
// 前序遍历二叉树，找到插入位置后直接插入即可
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		// 找到插入位置，进行插入操作
		root = &TreeNode{Val: val}
		return root
	}

	if root.Val > val {
		// 若当前节点的值大于目标值 -- 将目标值插入到当前节点的左子树中
		root.Left = insertIntoBST(root.Left, val)
	} else {
		// 若当前节点的值小于或等于目标值 -- 将目标值插入到当前节点的右子树中
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}
