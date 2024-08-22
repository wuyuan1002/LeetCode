package main

// 92. 反转链表 II

// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
// 请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

// reverseBetween .
// leetcode 206. 反转链表
//
// 先遍历链表寻找到左节点（过程中不进行反转操作），然后继续向后遍历寻找右节点（过程中进行反转操作）
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 构建一个虚拟头节点
	dummy := &ListNode{Next: head}

	// 先找到要反转范围左节点的前一个节点
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 继续向后遍历当前节点，直到找到右节点，过程中完成反转操作 -- leetcode 206. 反转链表
	node, next := pre.Next, pre.Next
	for i := 0; i < right-left; i++ {
		next = node.Next
		node.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummy.Next
}
