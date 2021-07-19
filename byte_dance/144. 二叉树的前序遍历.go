package main

// 144. 二叉树的前序遍历

// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

func main() {

}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	dfs11(root, &result)

	return result
}

func dfs11(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	*result = append(*result, node.Val)

	dfs11(node.Left, result)
	dfs11(node.Right, result)
}
