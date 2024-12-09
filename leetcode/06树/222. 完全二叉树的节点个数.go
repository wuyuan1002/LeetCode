package main

// 222. 完全二叉树的节点个数

// 给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
//
// 完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，
// 其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。
// 若最底层为第 h 层，则该层包含 1~ 2h 个节点。

// countNodes .
// 1. 递归遍历二叉树，过程中统计节点个数
// 2. 层级遍历二叉树，过程中统计节点个数
// 3. 利用完全二叉树的特点，使用满二叉树的公式求解：若一个满二叉树的深度为k，则其节点总数为(2^k)-1
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 返回左右子树的节点个数之和+1
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// countNodes222 利用完全二叉树的特点进行求解 -- 使用满二叉树的公式求解
func countNodes222(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth, rightDepth := 0, 0                // 当前节点的左右子树深度
	leftNode, rightNode := root.Left, root.Right // 要遍历的左右子节点

	// 分别统计当前节点的左右子树的深度
	for leftNode != nil {
		leftNode = leftNode.Left
		leftDepth++
	}
	for rightNode != nil {
		rightNode = rightNode.Right
		rightDepth++
	}

	// 由于给定的树是一棵完全二叉树，则
	// 若当前节点的左右子树深度相同，说明当前节点为一颗满二叉树的根，使用满二叉树公式求解当前节点子节点总数
	if leftDepth == rightDepth {
		return (2 << leftDepth) - 1
	}

	// 若当前节点不是一个满二叉树的根，则分别求其左右子树的节点个数+1
	return countNodes222(root.Left) + countNodes222(root.Right) + 1
}
