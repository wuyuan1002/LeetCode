package main

// 108. 将有序数组转换为二叉搜索树

// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
//
// 高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

// sortedArrayToBST .
//
// 相当于给了二叉搜索树的前序遍历序列，构建平衡的二叉搜索树
// 平衡二叉树的左右子树节点个数一致，所以序列的中间节点一定为根节点
// 所以先找到中点作为根节点, 前半部分为左子树，后半部分为右子树
func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	// 找到当前节点所在下标
	mid := len(nums) / 2

	// 构建当前节点并递归构建左右子节点
	node := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}

	return node
}
