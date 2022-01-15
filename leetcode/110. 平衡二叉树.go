package main

// 110. 平衡二叉树

// func main() {

// }

func isBalanced(root *TreeNode) bool {
	depth := 0
	return isBalance(root, &depth)
}

func isBalance(root *TreeNode, depth *int) bool {
	if root == nil {
		*depth = 0
		return true
	}

	lDepth, rDepth := 0, 0 // 左右子树深度
	lb := isBalance(root.Left, &lDepth)
	rb := isBalance(root.Right, &rDepth)

	diff := lDepth - rDepth
	if lb && rb && diff <= 1 && diff >= -1 {
		*depth = max(lDepth, rDepth) + 1
		return true
	}
	return false
}
