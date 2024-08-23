package main

// 206. 反转链表
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

// reverseList .
func reverseList(head *ListNode) *ListNode {
	// 前一个节点、当前节点、下一个节点
	pre, node, next := (*ListNode)(nil), head, head

	for node != nil {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}

	return pre
}
