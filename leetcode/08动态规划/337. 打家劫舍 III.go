package main

// 337. 打家劫舍 III

// 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
//
// 除了 root 之外，每栋房子有且只有一个“父“房子与之相连。
// 一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
// 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
//
// 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。

// rob337 .
// 偷了一个节点后，不可以再偷它的子节点
// 1. 先按层级广度优先遍历二叉树，算出每一层的总和，存在一个数组中，再和之前的打家劫舍一样不能偷取相邻的即可
// 2. 动态规划 -- 每个节点都可以选择偷或不偷
// 	    当前节点选择不偷: 当前节点能偷到的最大金额 = 左孩子能偷到的最大金额 + 右孩子能偷到的最大金额
// 	    当前节点选择偷: 当前节点能偷到的最大金额 = 左孩子选择自己不偷时能得到的金额 + 右孩子选择不偷时能得到的金额 + 当前节点的金额
//    总结果为root节点选择偷或不偷的最大值：max(选择偷root节点, 选择不偷root节点)
func rob337(root *TreeNode) int {
	// 计算root节点选择偷和不偷时能得到的金额数并返回两者的最大值
	rootRob, rootNoRob := dfsRob337(root)
	return max(rootRob, rootNoRob)
}

// dfsRob337 .
// 深度遍历给定节点node，返回选择偷当前节点和选择不偷当前节点所能得到的金额
func dfsRob337(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	// 计算当前节点的左右子节点分别选择偷和不偷所能得到的金额
	leftRob, leftNoRob := dfsRob337(node.Left)
	rightRob, rightNoRob := dfsRob337(node.Right)

	// 计算当前节点选择偷和不偷时能得到的金额数
	currentRob := leftNoRob + rightNoRob + node.Val
	currentNoRob := max(leftRob, leftNoRob) + max(rightRob, rightNoRob)

	return currentRob, currentNoRob
}

// TreeNode .
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
