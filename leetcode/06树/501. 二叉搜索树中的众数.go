package main

import "math"

// 501. 二叉搜索树中的众数

// 给你一个含重复值的二叉搜索树（BST）的根节点 root ，
// 找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。
// 如果树中有不止一个众数，可以按 任意顺序 返回。
//
// 假定 BST 满足如下定义：
//
// 结点左子树中所含节点的值 小于等于 当前节点的值
// 结点右子树中所含节点的值 大于等于 当前节点的值
// 左子树和右子树都是二叉搜索树

// findMode .
// 1. 普通二叉树：前中后序比那里二叉树，用map记录每个元素出现的次数，然后返回出现次数最多的元素
// 2. 二叉搜索树：中序遍历是递增的，所以相同元素都是相邻出现的，可以在中序遍历过程中统计出现次数最多的元素
func findMode(root *TreeNode) []int {
	count, maxCount := 0, 0
	preVal := math.MaxInt64
	result := make([]int, 0)
	dfsFindMode(root, &preVal, &count, &maxCount, &result)
	return result
}

// dfsFindMode 中序遍历二叉搜索树，统计出现次数最多的数字
// count: 当前节点数字出现的次数
// maxCount: 当前结果集中的众数出现的次数
// result: 总结果集
func dfsFindMode(node *TreeNode, preVal *int, count, maxCount *int, result *[]int) {
	if node == nil {
		return
	}

	// 遍历左子树
	dfsFindMode(node.Left, preVal, count, maxCount, result)

	if *preVal == math.MaxInt64 || node.Val != *preVal {
		// 当前节点node是遍历的整棵树的第一个节点或当前节点是一个新的数字 -- 更新当前节点出现的次数为1
		*count = 1
	} else {
		// 当前节点和前一个节点数字相同 -- 当前节点数字出现的次数++
		*count++
	}

	// 标记当前节点为已遍历
	*preVal = node.Val

	if *count == *maxCount {
		// 当前节点的值也是目前出现次数最多的值 -- 加入到结果集
		*result = append(*result, node.Val)
	} else if *count > *maxCount {
		// 当前节点的值出现次数最多，已经超过结果集中所记录的众数出现的次数
		// 更新目前众数的出现次数
		*maxCount = *count
		// 清空原来的众数结果，重新记录当前节点的值为众数
		*result = []int{node.Val}
	}

	// 遍历右子树
	dfsFindMode(node.Right, preVal, count, maxCount, result)
}
