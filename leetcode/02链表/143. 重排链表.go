package main

// 143. 重排链表

// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
// L0 → L1 → … → Ln - 1 → Ln
// 请将其重新排列后变为：
//
// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// reorderList .
// 获取链表中间节点并反转链表后半段，然后将后半段与前半段交替合并
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// 获取中间节点
	midNode := getMiddleNode(head)
	// 反转后半段链表
	tail := reverseList(midNode)

	// 合并链表前后两断
	l1, l2 := head, tail
	var l1Next, l2Next *ListNode
	for l1 != nil && l2 != nil {
		l1Next = l1.Next
		l2Next = l2.Next

		l1.Next = l2
		l1 = l1Next

		l2.Next = l1
		l2 = l2Next
	}
}

// getMiddleNode .
// 快慢指针
// 获取链表中间节点，快指针指向末尾时满指针正好指向中间节点
func getMiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
