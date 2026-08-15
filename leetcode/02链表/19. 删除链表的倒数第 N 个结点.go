package main

// 19. 删除链表的倒数第 N 个结点

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 进阶：你能尝试使用一趟扫描实现吗？

// removeNthFromEnd .
// Offer 22. 链表中倒数第k个节点
//
// 双指针
// 先让右指针向后移动n, 之后左右指针同时向后移动, 当右指针到末尾时, 要删除的元素正好时左指针的下一个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	l, r := dummy, head

	for i := 0; i < n; i++ {
		if r == nil {
			return head
		}
		r = r.Next
	}

	for r != nil {
		l = l.Next
		r = r.Next
	}

	l.Next = l.Next.Next

	return dummy.Next
}
