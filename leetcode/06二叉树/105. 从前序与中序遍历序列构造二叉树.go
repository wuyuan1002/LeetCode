package main

// 105. 从前序与中序遍历序列构造二叉树

// 给定两个整数数组 preorder 和 inorder ，
// 其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，
// 请构造二叉树并返回其根节点。

// buildTree105 .
// 同 leetcode 106. 从中序与后序遍历序列构造二叉树
// 前序遍历为根左右、中序遍历为左根右、后序遍历为左右根
func buildTree105(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 || inorder == nil || len(inorder) == 0 {
		return nil
	}

	// 当前节点的值
	currentVal := preorder[0]

	// 查找当前节点在中序遍历中的下标
	index := 0
	for index < len(inorder) && inorder[index] != currentVal {
		index++
	}

	// 切分左右子树的中序遍历和后序遍历列表
	linorder := inorder[:index]
	rinorder := inorder[index+1:]
	lpreorder := preorder[1 : index+1]
	rpreorder := preorder[index+1:]

	// 构造当前节点并递归构造左右子树
	node := &TreeNode{
		Val:   currentVal,
		Left:  buildTree105(lpreorder, linorder),
		Right: buildTree105(rpreorder, rinorder),
	}

	return node
}
