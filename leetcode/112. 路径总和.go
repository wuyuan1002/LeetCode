package main

// 112. 路径总和

func main() {

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs12(root, 0, targetSum)
}

func dfs12(node *TreeNode, currentSum, target int) bool {
	if node == nil {
		return false
	}

	currentSum += node.Val
	if node.Left == nil && node.Right == nil && currentSum == target {
		return true
	}

	return dfs12(node.Left, currentSum, target) || dfs12(node.Right, currentSum, target)
}
