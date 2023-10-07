package main

// 437. 路径总和 III

// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，
// 求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

// pathSum437 .
// 同 leetcode 112. 路径总和、leetcode 113. 路径总和 II、leetcode 257. 二叉树的所有路径
// 回溯 -- 前序递归遍历二叉树，遍历过程中统计满足条件的路径数量
func pathSum437(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	result := 0
	// 计算从跟节点出发的满足条件的路径条数
	dfsPathSum437(root, 0, targetSum, &result)

	// 计算以左节点出发的满足条件的路径条数
	result += pathSum437(root.Left, targetSum)
	// 计算以右节点出发的满足条件的路径条数
	result += pathSum437(root.Right, targetSum)

	return result
}

// dfsPathSum437 统计以给定节点node为根节点出发的满足条件的路径个数
// 此函数功能类似于 leetcode 113. 路径总和 II
func dfsPathSum437(node *TreeNode, currentSum, targetSum int, result *int) {
	if node == nil {
		return
	}

	// 将当前节点的值加到当前和上
	currentSum += node.Val

	// 若当前节点满足条件，更新总结果
	if currentSum == targetSum {
		*result++
	}

	// 继续递归当前节点的左右子节点
	dfsPathSum437(node.Left, currentSum, targetSum, result)
	dfsPathSum437(node.Right, currentSum, targetSum, result)
}
