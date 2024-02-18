package main

// 109. 有序链表转换二叉搜索树

// 给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为高度平衡的二叉搜索树。
//
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差不超过 1。

// globalHead 用于指向当前正在构建的节点指针（当前的链表头节点），全局在整个链表上向后移动
var globalHead *ListNode

// sortedListToBST .
// 1. 将链表中序遍历后得到有序数组，然后同 leetcode 108 进行解答
// 2. 使用快慢指针法找到链表中间节点，此为当前树的根节点，然后递归构建其左子树和右子树
// 3. 不先找到根节点再分别构建左右子树，而是使用中序遍历，先构建左子树再构建根节点再构建右子树，构建完左子树后自然会得到根节点是什么，不过需要提前确定树的大小（因为树是高度平衡的，根节点一定是中间节点）
func sortedListToBST(head *ListNode) *TreeNode {
	// 当前正在遍历的是第一个节点
	globalHead = head

	// 求出链表长度 -- 整棵树的长度，也就是当前要求根节点的最右节点的下标
	length := 0
	for ; head != nil; head = head.Next {
		length += 1
	}

	// 构建第一个根节点及其左右子树
	return buildTree(0, length-1)
}

// buildTree 构建平衡二叉搜索树
// left、right: 对于当前根节点的左右子节点在链表中的下标（左闭右闭）
func buildTree(left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// 计算当前根节点所处下标 -- 即根节点下标
	mid := left + (right-left)/2

	// 当前当前节点 -- 即根节点
	root := &TreeNode{}

	// 构建当前节点的左子树
	root.Left = buildTree(left, mid-1)

	// 构建当前节点，并将正在构造的节点指针向后移动一位 -- 即根节点
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
