package main

// 剑指 Offer 55 - II. 平衡二叉树

// 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。
// 如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

// func main() {

// }

// 先求出左右子树的深度，比较是不是平衡，再分别求左右子树是不是平衡 -- 从上到下，先序遍历，先看当前节点是不是，再看子节点是不是
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	diff := left - right
	if diff < -1 || diff > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func isBalanced1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	depth := 0 // 根节点深度
	return isBalance(root, &depth)
}

// 自下而上，后序遍历 -- 先判断左右子树是不是平衡，再判断当前节点是不是平衡
func isBalance(root *TreeNode, depth *int) bool { // depth 表示当前树的深度
	if root == nil {
		*depth = 0 // 空树的深度为0
		return true
	}

	leftDepth, rightDepth := 0, 0 // 表示左右子树的深度

	lb := isBalance(root.Left, &leftDepth)   // 左节点是不是平衡
	rb := isBalance(root.Right, &rightDepth) // 右节点是不是平衡

	// 判断当前节点是不是平衡的 -- 左右子树都是平衡的并且当前节点的左右子树深度差小于1
	diff := leftDepth - rightDepth
	if lb && rb && diff >= -1 && diff <= 1 { // 如果左右子树都是平衡的，并且当前节点的左右子树深度差小于1，说明当前节点是平衡的
		*depth = max(leftDepth, rightDepth) + 1 // 当前节点的深度
		return true
	}

	return false
}

// 第二次做
// 自下而上，后序遍历 -- 先判断左右子树是不是平衡，再判断当前树是不是平衡
func isBalance1(root *TreeNode, depth *int) bool {
	if root == nil {
		*depth = 0 // 空树的深度为0
		return true
	}

	leftDepth, rightDepth := 0, 0                                                 // 左右子树的深度
	if isBalance1(root.Left, &leftDepth) && isBalance1(root.Right, &rightDepth) { // 若左右子树都是平衡的
		diff := leftDepth - rightDepth
		if diff >= -1 && diff <= 1 { // 若左右子树深度差小于1
			*depth = max(leftDepth, rightDepth) + 1 // 当前节点的深度
			return true
		}
	}
	return false
}
