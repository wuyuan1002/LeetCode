package main

// 206. 反转链表

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

// func main() {

// }

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre, node, next *ListNode = nil, head, head
	for node != nil {
		next = node.Next

		node.Next = pre

		pre = node
		node = next
	}

	return pre
}

// ---------
// 第二次做
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre, node, next *ListNode = nil, head, head
	for node != nil {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}

type ListNode struct {
	Val  int
	Next *ListNode
}
