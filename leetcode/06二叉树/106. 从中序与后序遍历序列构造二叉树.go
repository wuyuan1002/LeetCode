package main

// 106. 从中序与后序遍历序列构造二叉树

// 给定两个整数数组 inorder 和 postorder ，
// 其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，
// 请你构造并返回这颗 二叉树 。

// buildTree .
// 同 leetcode 105. 从前序与中序遍历序列构造二叉树
// 前序遍历为根左右、中序遍历为左根右、后序遍历为左右根
func buildTree106(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || len(inorder) == 0 || postorder == nil || len(postorder) == 0 {
		return nil
	}

	// 当前节点的值
	currentVal := postorder[len(postorder)-1]

	// 查找当前节点在中序遍历中的下标
	index := 0
	for index < len(inorder) && inorder[index] != currentVal {
		index++
	}

	// 切分左右子树的中序遍历和后序遍历列表
	linorder := inorder[:index]
	rinorder := inorder[index+1:]
	lpostorder := postorder[:index]
	rpostorder := postorder[index : len(postorder)-1]

	// 构造当前节点并递归构造左右子树
	node := &TreeNode{
		Val:   currentVal,
		Left:  buildTree106(linorder, lpostorder),
		Right: buildTree106(rinorder, rpostorder),
	}

	return node
}
