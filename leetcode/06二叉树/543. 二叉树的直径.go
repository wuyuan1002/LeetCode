package main

// 543. 二叉树的直径

// 给你一棵二叉树的根节点，返回该树的 直径 。
//
// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
//
// 两节点之间路径的 长度 由它们之间边数表示。

// diameterOfBinaryTree .
// 同 leetcode 104. 二叉树的最大深度
// 计算二叉树的最大深度，过程中更新二叉树的最大直径，直径为某个节点的左右子树的最大深度之和+1
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 1 // 最大直径
	maxDepthDiameter(root, &maxDiameter)

	// 由于直径为两节点间的边数而不是节点数，而我们前面求的最大深度都是求的节点数
	// 所以要把直径默认值置为1，返回时将节点数-1即为为节点间的边数
	return maxDiameter - 1
}

// maxDepthDiameter 求二叉树给定节点的最大深度，同时更新二叉树的最大直径
func maxDepthDiameter(node *TreeNode, maxDiameter *int) int {
	if node == nil {
		return 0
	}

	// 分别求出左右子树的最大深度
	leftDepth := maxDepthDiameter(node.Left, maxDiameter)
	rightDepth := maxDepthDiameter(node.Right, maxDiameter)

	// 更新二叉树的最大直径 -- 经过当前节点的直径与原最大直径的最大值
	*maxDiameter = max(*maxDiameter, leftDepth+rightDepth+1)

	// 返回当前节点的最大深度+1 -- 当前节点的深度也需要被加上
	return max(leftDepth, rightDepth) + 1
}
