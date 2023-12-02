package main

// 144. 二叉树的前序遍历

// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

// preorderTraversal 前序遍历二叉树
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	dfsPreOrder(root, &result)
	return result
}

// dfsPreOrder 深度优先前序遍历二叉树
func dfsPreOrder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	*result = append(*result, node.Val)
	dfsPreOrder(node.Left, result)
	dfsPreOrder(node.Right, result)
}

// TreeNode .
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
