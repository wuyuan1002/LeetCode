package main

// 1008. 前序遍历构造二叉搜索树

// 给定一个整数数组，它表示BST(即 二叉搜索树 )的 先序遍历 ，构造树并返回其根。
//
// 保证 对于给定的测试用例，总是有可能找到具有给定需求的二叉搜索树。
//
// 二叉搜索树 是一棵二叉树，其中每个节点， Node.left 的任何后代的值 严格小于 Node.val , Node.right 的任何后代的值 严格大于 Node.val。
// 二叉树的 前序遍历 首先显示节点的值，然后遍历Node.left，最后遍历Node.right。

// bstFromPreorder .
// 1. 二叉搜索树的中序遍历是有序的，因此可以将前序遍历数组进行排序从而得到中序遍历数组，此题便转化为 leetcode 105. 从前序与中序遍历序列构造二叉树
// 2. 使用二分查找寻找一个节点的左右子树分界点，然后将数组剩余根据这个分界点分割为其左右子树节点，进行递归构造左右子树 -- 左子树的值小于根节点，右子树的值大于根节点
func bstFromPreorder(preorder []int) *TreeNode {
	return dfsFstFromPreorder(preorder, 0, len(preorder)-1)
}

// dfsFstFromPreorder 进行递归构造给定范围节点，left、right表示本次构造的节点选取范围
func dfsFstFromPreorder(preorder []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// 构造本次的根节点 -- 前序遍历序列的第一个节点为根节点
	root := &TreeNode{Val: preorder[left]}

	// 若当前区间只有根节点一个节点，根节点没有左右子树，则直接返回根节点
	if left == right {
		return root
	}

	// 二分查找当前根节点的左右子树节点分界点 -- 查找结束后，l指向的是左边界的最后一个值，左右子树节点范围分别为 [left+1, l]、[l+1, right]
	// 在区间 [left, right] 里找最后一个小于 preorder[left] 的下标
	// 注意这里设置区间的左边界为 left ，不能是 left + 1
	// 这是因为考虑到区间只有 2 个元素 [left, right] 的情况，第 1 个部分为空区间，第 2 部分只有一个元素 right
	l, r := left, right
	for l < r {
		mid := l + (r-l+1)/2
		if preorder[mid] < preorder[left] {
			l = mid
		} else {
			r = mid - 1
		}
	}

	// 递归构造当前节点的左右子树
	root.Left = dfsFstFromPreorder(preorder, left+1, l)
	root.Right = dfsFstFromPreorder(preorder, l+1, right)

	return root
}
