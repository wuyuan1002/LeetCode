package main

// 25. K个一组翻转链表

// 给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
// k是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 进阶：
// 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

// reverseKGroup .
// leetcode 24. 两两交换链表中的节点
//
// 按照k个节点进行分组，对每一组进行反转链表，
// 然后将反转后当前组的尾节点指向反转后下一组的头节点，最终返回当前组的头节点
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	// 分割前k个节点，寻找到下一个组的起始节点
	nextGroupHead := head
	for i := 0; i < k; i++ {
		if nextGroupHead == nil {
			return head
		}
		nextGroupHead = nextGroupHead.Next
	}

	// 反转当前组
	newHead := reverse(head, nextGroupHead)
	// 翻转之后的节点，并将反转后当前组的尾节点指向反转后下一组的头节点
	head.Next = reverseKGroup(nextGroupHead, k)

	return newHead
}

// reverse 反转链表 -- 反转两个节点之间的节点 [)
func reverse(head, nextGroupHead *ListNode) *ListNode {
	pre, node, next := (*ListNode)(nil), head, head
	for node != nextGroupHead {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}
