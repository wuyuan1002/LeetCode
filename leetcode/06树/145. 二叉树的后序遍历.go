package main

// 145. 二叉树的后序遍历

// 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。

// postorderTraversal 递归遍历二叉树 -- 后序遍历
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	dfsPostOrder(root, &result)
	return result
}

// dfsPostOrder 深度优先后序遍历二叉树
func dfsPostOrder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	dfsPostOrder(node.Left, result)
	dfsPostOrder(node.Right, result)
	*result = append(*result, node.Val)
}
