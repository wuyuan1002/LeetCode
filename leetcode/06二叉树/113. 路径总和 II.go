package main

// 113. 路径总和 II

// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，
// 找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
// 叶子节点 是指没有子节点的节点。

// pathSum113 .
// 同 leetcode 112. 路径总和、leetcode 437. 路径总和 III、leetcode 257. 二叉树的所有路径
// 回溯 -- 前序递归遍历二叉树，遍历过程中记录满足条件的路径
func pathSum113(root *TreeNode, targetSum int) [][]int {
	res := make([]int, 0)      // 一次遍历路径上的节点
	result := make([][]int, 0) // 总结果

	dfsPathSum113(root, 0, targetSum, &res, &result)
	return result
}

// dfsPathSum 回溯递归遍历节点获取路径
func dfsPathSum113(node *TreeNode, currentSum, targetSum int, res *[]int, result *[][]int) {
	if node == nil {
		return
	}

	// 更新当前路径上目前的和
	currentSum += node.Val

	// 将当前节点的值加入到本次遍历的路径中
	*res = append(*res, node.Val)

	if node.Left == nil && node.Right == nil && currentSum == targetSum {
		// 若当前节点为叶子节点且当前和等于目标值 -- 记录一条路径总和到总结果
		temp := make([]int, len(*res))
		copy(temp, *res)
		*result = append(*result, temp)
	} else {
		// 若当前节点不满足条件 -- 继续遍历当前节点的左右子节点
		dfsPathSum113(node.Left, currentSum, targetSum, res, result)
		dfsPathSum113(node.Right, currentSum, targetSum, res, result)
	}

	// 当前节点遍历完成，将当前节点的值移出本次遍历的路径中
	*res = (*res)[:len(*res)-1]
}
