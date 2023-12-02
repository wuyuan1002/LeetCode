package main

// 110. 平衡二叉树

// 给定一个二叉树，判断它是否是高度平衡的二叉树。
//
// 本题中，一棵高度平衡二叉树定义为：
// 一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1 。

// isBalanced .
// 同 Offer 55-2、leetcode 104. 二叉树的最大深度
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

	leftDepth, rightDepth := 0, 0            // 左右子节点的最大深度
	lb := isBalance(node.Left, &leftDepth)   // 左子树是否平衡
	rb := isBalance(node.Right, &rightDepth) // 右子树是否平衡

	diff := leftDepth - rightDepth // 当前节点的左右子树高度差

	// 当前节点平衡
	if lb && rb && diff >= -1 && diff <= 1 {
		// 获取当前节点的最大深度并返回
		*maxDepth = max(leftDepth, rightDepth) + 1
		return true
	}

	// 当前节点不平衡
	return false
}
