package main

// 110. 平衡二叉树

// 给定一个二叉树，判断它是否是高度平衡的二叉树。
//
// 本题中，一棵高度平衡二叉树定义为：
// 一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1 。

// func main() {

// }

// offer 55-2
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	depth := 0
	return isBalance(root, &depth)
}

// isBalance 求一颗子树是否为平衡二叉树
// depth: 当前树的深度
func isBalance(node *TreeNode, depth *int) bool {
	if node == nil {
		*depth = 0
		return true
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	leftDepth, rightDepth := 0, 0            // 左右子树深度
	lb := isBalance(node.Left, &leftDepth)   // 左树是否为平衡二叉树
	rb := isBalance(node.Right, &rightDepth) // 右树是否为平衡二叉树

	diff := leftDepth - rightDepth // 左右子树高度差
	if lb && rb && diff >= -1 && diff <= 1 {
		*depth = max(leftDepth, rightDepth) + 1
		return true
	}
	return false
}
