package main

// 543. 二叉树的直径

// 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。
// 这条路径可能穿过也可能不穿过根结点。

// func main() {

// }

// 后序遍历二叉树 -- 分别求每个节点左右子树的深度，求每个节点左右子树深度和的最大值
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	diamter := 1 // 最大路线长度
	_ = dfs6(root, &diamter)
	return diamter - 1
}

func dfs6(node *TreeNode, diamter *int) int {
	if node == nil {
		return 0
	}

	leftDepth := dfs6(node.Left, diamter)
	rightDepth := dfs6(node.Right, diamter)

	*diamter = max(*diamter, leftDepth+rightDepth+1)

	return max(leftDepth, rightDepth) + 1
}
