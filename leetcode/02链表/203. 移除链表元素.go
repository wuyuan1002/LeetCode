package main

// 203. 移除链表元素

// 给你一个链表的头节点 head 和一个整数 val ，
// 请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

// removeElements .
// leetcode 82. 删除排序链表中的重复元素 II
// leetcode 83. 删除排序链表中的重复元素
//
// 构建一个虚拟头节点（防止head正好也是要被删除的节点），然后不断遍历链表节点，将要删除的节点从链表中移除
func removeElements(head *ListNode, val int) *ListNode {

	// 虚拟头节点
	dummy := &ListNode{Next: head}

	// 不断遍历链表节点，若节点需要被删除则将该节点从链表中移除
	for node := dummy; node != nil && node.Next != nil; {
		if node.Next.Val == val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

	return dummy.Next
}

// ListNode .
type ListNode struct {
	Val  int
	Next *ListNode
}
