package main

// 95. 不同的二叉搜索树 II

// 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。

// generateTrees .
// 同 leetcode 96. 不同的二叉搜索树 -- 96提只需返回可能的情况个数，本题要返回所有可能的解
// 本题不能使用动态规划，动态规划只能得出解的个数而得不到解具体是什么，因此需要使用递归真正生成每棵树
// 从1-n中挨个选择每个节点i作为最终树的跟节点，然后分别递归求1-i和i-n的所有不同二叉树，它们求出的就是当前节点i的所有子树的所有可能，然后从其中获取每个子树拼接到当前树上即得到一个解
func generateTrees(n int) []*TreeNode {
	// 生成根节点为1-n范围的所有解
	return generate(1, n)
}

// generate 用于生成给定范围[start, end]的不同二叉搜索树的所有解
func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	// 节点范围为[start, end]的所有解
	allTrees := make([]*TreeNode, 0)

	// 遍历[start, end] -- 从中选择每一个节点i作为跟节点，并分别生成所有子树的不同二叉树的解
	for i := start; i <= end; i++ {
		// 分别求解左右子树的所有解
		leftTrees := generate(start, i-1)
		rightTrees := generate(i+1, end)

		// 从左右子树解中分别选择左右子树，拼接到当前根节点上，得到以i为根节点的一个解
		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				// 构建当前节点并添加到总结果中
				currentTree := &TreeNode{Val: i, Left: leftTree, Right: rightTree}
				allTrees = append(allTrees, currentTree)
			}
		}
	}

	return allTrees
}
