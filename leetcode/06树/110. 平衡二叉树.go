package main

// 110. 平衡二叉树

// 给定一个二叉树，判断它是否是高度平衡的二叉树。
//
// 本题中，一棵高度平衡二叉树定义为：
// 一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1 。

// isBalanced .
// leetcode 104. 二叉树的最大深度
// Offer 55 - II. 平衡二叉树
//
// 递归求出左右子树的最大深度，然后判断该节点是否平衡
func isBalanced(root *TreeNode) bool {
	depth := 0
	return isBalance(root, &depth)
}

// isBalance 求一颗树是否为平衡二叉树
// maxDepth: 当前节点的最大深度
func isBalance(node *TreeNode, maxDepth *int) bool {
	if node == nil {
		*maxDepth = 0
		return true
	}

	leftDepth, rightDepth := 0, 0                        // 左右子节点的最大深度
	isLeftBalance := isBalance(node.Left, &leftDepth)    // 左子树是否平衡
	isRightBalance := isBalance(node.Right, &rightDepth) // 右子树是否平衡

	diff := leftDepth - rightDepth // 当前节点的左右子树高度差

	// 若当前节点的左右子树都平衡并且左右子树的高度差不超过1，则说明当前节点是平衡的
	if isLeftBalance && isRightBalance && diff >= -1 && diff <= 1 {
		// 获取当前节点的最大深度并返回
		*maxDepth = max(leftDepth, rightDepth) + 1
		return true
	}

	// 当前节点不平衡
	return false
}
