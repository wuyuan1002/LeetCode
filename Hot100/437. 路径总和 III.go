package main

// 437. 路径总和 III

// 给定一个二叉树，它的每个结点都存放着一个整数值。
// 找出路径和等于给定数值的路径总数。
// 路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
// 二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

func main() {

}

// 回溯法
func pathSum1(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return dfs11(root, targetSum) + pathSum1(root.Left, targetSum) + pathSum1(root.Right, targetSum)
}

func dfs11(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// 从当前节点出发的路径个数
	sum := 0
	if root.Val == targetSum {
		sum++
	}
	sum += dfs11(root.Left, targetSum-root.Val)
	sum += dfs11(root.Right, targetSum-root.Val)
	return sum
}

// func pathSum1(root *TreeNode, targetSum int) int {
// 	if root == nil {
// 		return 0
// 	}
//
// 	result := 0
// 	dfs11(root, 0, targetSum, &result)
// 	return result
// }
//
// func dfs11(root *TreeNode, currentSum, targetSum int, result *int) {
// 	if root == nil {
// 		return
// 	}
//
// 	// 计算加上当前节点后的总和
// 	sum := root.Val + currentSum
// 	if sum == targetSum {
// 		*result++
// 	}
//
// 	if sum <= targetSum {
// 		dfs11(root.Left, sum, targetSum, result)
// 		dfs11(root.Right, sum, targetSum, result)
// 	}
// 	dfs11(root.Left, root.Val, targetSum, result)
// 	dfs11(root.Right, root.Val, targetSum, result)
// }
