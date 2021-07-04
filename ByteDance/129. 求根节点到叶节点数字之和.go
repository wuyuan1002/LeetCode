package main

import (
	"strconv"
)

// 129. 求根节点到叶节点数字之和

// 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
// 每条从根节点到叶节点的路径都代表一个数字：
//
// 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
// 计算从根节点到叶节点生成的 所有数字之和 。
//
// 叶节点 是指没有子节点的节点。

func main() {

}

func sumNumbers(root *TreeNode) int {
	result := 0
	nums := make([]byte, 0)
	dfs4(root, &nums, &result)
	return result
}

// nums: 存放路径上的数字
func dfs4(node *TreeNode, nums *[]byte, result *int) {
	if node == nil {
		return
	}

	// 把当前节点的值放入路径中
	*nums = append(*nums, byte(node.Val+'0'))

	if node.Left == nil && node.Right == nil {
		// 若当前节点是叶节点，计算总和
		n, _ := strconv.Atoi(string(*nums))
		*result += n
	} else {
		// 若不是叶节点，继续向下遍历
		dfs4(node.Left, nums, result)
		dfs4(node.Right, nums, result)
	}

	// 将当前节点移出路径
	*nums = (*nums)[:len(*nums)-1]
}
