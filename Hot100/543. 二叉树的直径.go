package main

// 543. 二叉树的直径

// 给定一棵二叉树，你需要计算它的直径长度。
// 一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

func main() {

}

// 中序遍历二叉树
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth = 1
	maxDepth1(root)

	return depth - 1
}

var depth = 0 // 最大路线长度

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth1(root.Left)
	rightDepth := maxDepth1(root.Right)

	// 更新最大路线长度
	depth = max(depth, leftDepth+rightDepth+1)

	// 返回当前节点的最大深度
	return max(leftDepth, rightDepth) + 1
}
