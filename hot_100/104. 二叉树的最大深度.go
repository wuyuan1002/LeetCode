package main

// 104. 二叉树的最大深度

// 给定一个二叉树，找出其最大深度。
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
// 说明: 叶子节点是指没有子节点的节点。

// func main() {

// }

// 同offer 55-1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	// 分别求出左右子树深度
	leftDeph := maxDepth(root.Left)
	rightDeph := maxDepth(root.Right)

	// 返回大的那个+1
	return max(leftDeph, rightDeph) + 1
}
