package main

// 206. 反转链表

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

// func main() {

// }

// 同offer 24
// 三指针
func reverseList(head *ListNode) *ListNode {
	pre, node, next := (*ListNode)(nil), head, head
	for node != nil {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}

// 第二次做
func reverseList1(head *ListNode) *ListNode {
	pre, node, next := (*ListNode)(nil), head, head
	for node != nil {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}
