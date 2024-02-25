package main

import "strconv"

// 129. 求根节点到叶节点数字之和

// 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
// 每条从根节点到叶节点的路径都代表一个数字：
//
// 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
// 计算从根节点到叶节点生成的 所有数字之和 。
//
// 叶节点 是指没有子节点的节点。

// sumNumbers .
// 回溯法 -- 遍历到叶节点则计算总和，中间节点则继续向下遍历
func sumNumbers(root *TreeNode) int {
	result := 0
	nums := make([]byte, 0)
	dfsSumNumbers(root, &nums, &result)
	return result
}

// dfsSumNumbers 遍历给定的节点是否为叶节点，若是则计算并相加到总结果，若不是则继续递归其左右子树
// node: 当前节点
// nums: 一条回溯路径
// result: 总结果
func dfsSumNumbers(node *TreeNode, nums *[]byte, result *int) {
	if node == nil {
		return
	}

	// 将当前节点加入到回溯路径
	*nums = append(*nums, byte(node.Val+'0'))

	if node.Left == nil && node.Right == nil {
		// 若当前节点是叶节点 -- 计算总和并加到总结果
		n, _ := strconv.Atoi(string(*nums))
		*result += n
	} else {
		// 若当前节点不是叶节点 -- 继续递归其左右子树
		dfsSumNumbers(node.Left, nums, result)
		dfsSumNumbers(node.Right, nums, result)
	}

	// 将当前节点移出回溯路径
	*nums = (*nums)[:len(*nums)-1]
}

// TreeNode .
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
