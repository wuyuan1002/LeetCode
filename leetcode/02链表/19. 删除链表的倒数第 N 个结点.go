package main

// 19. 删除链表的倒数第 N 个结点

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 进阶：你能尝试使用一趟扫描实现吗？

// removeNthFromEnd .
// 同Offer 22
// 先让右指针向后移动n, 之后左右指针同时向后移动, 当右指针到末尾时, 要删除的元素正好时左指针的下一个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	l, r := head, head

	// 右指针先向后移动n
	for i := 0; i < n; i++ {
		if r == nil {
			return head
		}
		r = r.Next
	}

	// 倒数第n个节点恰巧是头节点 -- 要删除的元素正好的head
	if r == nil {
		return head.Next
	}

	// 双指针同时向后移动
	for r.Next != nil {
		l = l.Next
		r = r.Next
	}

	// 删除倒数第n个节点
	l.Next = l.Next.Next
	return head

}
