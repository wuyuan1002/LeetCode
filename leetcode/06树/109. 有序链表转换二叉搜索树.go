package main

// 109. 有序链表转换二叉搜索树

// 给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为高度平衡的二叉搜索树。
//
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差不超过 1。

// globalHead 用于指向当前正在构建的节点指针，全局在整个链表上向后移动
var globalHead *ListNode

// sortedListToBST .
// leetcode 108. 将有序数组转换为二叉搜索树
//
// *** 由于最终的树是平衡二叉树，所以根节点一定是链表的中间节点 ***
//
// 1. 将有序链表转换为有序数组，然后使用 leetcode 108 的方法进行解答
// 2. 使用快慢指针法找到链表中间节点，此为当前树的根节点，然后递归构建其左子树和右子树
// 3. 不先找到根节点再分别构建左右子树，而是类似中序遍历，先构建左子树再构建根节点再构建右子树，
//    构建完左子树后自然会得到根节点是什么，不过需要提前确定树的大小
func sortedListToBST(head *ListNode) *TreeNode {
	// 从链表第一个节点开始向后挨个遍历进行构造
	globalHead = head

	// 求出链表长度 -- 整棵树的长度，也就是当前根节点的最右子节点的下标
	length := 0
	for ; head != nil; head = head.Next {
		length += 1
	}

	// 构建第一个根节点及其左右子树
	return buildTree(0, length-1)
}

// buildTree 构建平衡二叉搜索树
// left、right: 对于当前根节点的左右子节点在链表中的下标（左闭右闭 -- [left, right]）
func buildTree(left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// 当前根节点下标
	mid := left + (right-left)/2

	// 先构建一个空的当前节点
	root := &TreeNode{}

	// 构建当前节点的左子树
	root.Left = buildTree(left, mid-1)

	// 构建当前根节点，并将正在构造的节点指针globalHead向后移动一位
	// globalHead表示当前正在构造的节点，在链表上从前向后移动，
	// 在前面构造左子树时，它不断向后移动进行构造，到此处时mid前的节点都已构造完成
	root.Val = globalHead.Val
	globalHead = globalHead.Next

	// 构建当前节点的右子树
	root.Right = buildTree(mid+1, right)

	return root
}

// ListNode .
type ListNode struct {
	Val  int
	Next *ListNode
}
