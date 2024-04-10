package main

// 515. 在每个树行中找最大值

// 给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。

// largestValues .
// 1. 广度优先搜索 -- 层序遍历二叉树，遍历一层的同时计算当前层的最大值
// 2. 深度优先搜索 -- 前序遍历二叉树，总结果的下标表示当前节点的层数，对应的值为当前遍历到的该层的最大值，遍历过程中每遍历到一个节点都与当前节点所在层的最大值进行比较并替换
func largestValues(root *TreeNode) []int {
	result := make([]int, 0)
	dfsLargestValues(root, 0, &result)
	return result
}

// dfsPreOrder 深度优先前序遍历二叉树进行每层最大值的统计
// node: 当前节点
// currentLevel: 当前节点所在层
// result: 总结果
func dfsLargestValues(node *TreeNode, currentLevel int, result *[]int) {
	if node == nil {
		return
	}

	// 遍历当前节点 -- 若当前节点是新的层则将当前节点的值加入到总结果，若当前节点所在层已被遍历过则比较并替换当前节点所在层的最大值
	if currentLevel == len(*result) {
		*result = append(*result, node.Val)
	} else {
		(*result)[currentLevel] = max((*result)[currentLevel], node.Val)
	}

	// 遍历当前节点的左右节点，子节点一定在当前节点的下一层，所以节点所在层数+1
	dfsLargestValues(node.Left, currentLevel+1, result)
	dfsLargestValues(node.Right, currentLevel+1, result)
}
