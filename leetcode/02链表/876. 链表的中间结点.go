package main

// 876. 链表的中间结点

// 给你单链表的头结点 head ，请你找出并返回链表的中间结点。
//
// 如果有两个中间结点，则返回第二个中间结点。

// middleNode .
// 快慢指针
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
