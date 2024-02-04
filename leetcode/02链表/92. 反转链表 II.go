package main

// 92. 反转链表 II

// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

// reverseBetween .
// 1. 使用 leetcode 206. 反转链表，先遍历链表找到需要反转的节点范围，然后反转链表 -- 共需要遍历链表两次
// 2. 寻找到左节点后继续向后遍历寻找右节点，过程中完成反转操作 -- 总共需要遍历一次
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 构建一个虚拟头节点
	dummy := &ListNode{Next: head}

	// 先找到要反转范围左节点的前一个节点
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 继续向后遍历寻找右节点，过程中完成反转操作 -- 不断将右侧节点移动到pre的next
	node, next := pre.Next, pre.Next
	for i := 0; i < right-left; i++ {
		next = node.Next
		node.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummy.Next
}
