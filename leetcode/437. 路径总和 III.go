package main

// 437. 路径总和 III

// 给定一个二叉树的根节点 root，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

func main() {

}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	result := 0
	dfs21(root, targetSum, 0, &result)
	result += pathSum(root.Left, targetSum)
	result += pathSum(root.Right, targetSum)
	return result
}

func dfs21(node *TreeNode, targetSum, currentSum int, result *int) {
	if node == nil {
		return
	}

	sum := currentSum + node.Val
	if sum == targetSum {
		*result++
	}

	dfs21(node.Left, targetSum, sum, result)
	dfs21(node.Right, targetSum, sum, result)
}
