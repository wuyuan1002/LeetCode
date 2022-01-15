package main

// 111. 二叉树的最小深度

// 给定一个二叉树，找出其最小深度。
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
// 说明：叶子节点是指没有子节点的节点。

// func main() {

// }

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
