package main

// 112. 路径总和

// 给你二叉树的根节点root 和一个表示目标和的整数targetSum ，
// 判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和targetSum 。
//
// 叶子节点是指没有子节点的节点。

func main() {

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs2(root, 0, targetSum)
}

func dfs2(node *TreeNode, currentSum, targetSum int) bool {
	if node == nil {
		return false
	}

	currentSum += node.Val
	if currentSum == targetSum && node.Left == nil && node.Right == nil {
		return true
	}

	return dfs2(node.Left, currentSum, targetSum) || dfs2(node.Right, currentSum, targetSum)
}
