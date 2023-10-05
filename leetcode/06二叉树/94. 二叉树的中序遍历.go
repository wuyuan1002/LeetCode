package main

// 94. 二叉树的中序遍历

// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

// inorderTraversal 中序遍历二叉树
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	dfsInOrder(root, &result)
	return result
}

// dfsInOrder 深度优先中序遍历二叉树
func dfsInOrder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	dfsInOrder(node.Left, result)
	*result = append(*result, node.Val)
	dfsInOrder(node.Right, result)
}
