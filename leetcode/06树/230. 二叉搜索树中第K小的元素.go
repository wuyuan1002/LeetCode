package main

// 230. 二叉搜索树中第K小的元素

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

// kthSmallest .
//
// 二叉搜索树的中序遍历是递增的，所以中序遍历二叉树，在遍历过程中查找第k小的数字
func kthSmallest(root *TreeNode, k int) int {
	result := 0
	dfsKthSmallest(root, &k, &result)
	return result
}

// dfsKthSmallest 中序遍历二叉树寻找第k小的数字，找到后直接返回，不继续递归
func dfsKthSmallest(node *TreeNode, k *int, result *int) {
	// 若当前节点为nil或已经找到结果则直接返回
	if node == nil || *k <= 0 {
		return
	}

	dfsKthSmallest(node.Left, k, result)

	// 递减计数，若当前数字正好是第k小则给结果赋值
	if *k--; *k == 0 {
		*result = node.Val
	}

	dfsKthSmallest(node.Right, k, result)
}
