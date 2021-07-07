package main

// 94. 二叉树的中序遍历

// 给定一个二叉树的根节点 root ，返回它的 中序 遍历。

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	inorder(root, &result)
	return result
}

func inorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	inorder(node.Left, result)
	*result = append(*result, node.Val)
	inorder(node.Right, result)
}
