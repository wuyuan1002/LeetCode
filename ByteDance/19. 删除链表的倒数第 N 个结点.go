package main

// 19. 删除链表的倒数第N个结点

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 进阶：你能尝试使用一趟扫描实现吗？

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	l, r := head, head
	for i := 0; i < n; i++ {
		if r == nil {
			return head
		}
		r = r.Next
	}

	for r != nil && r.Next != nil {
		l = l.Next
		r = r.Next
	}

	if r == nil {
		return head.Next
	} else {
		l.Next = l.Next.Next
		return head
	}
}
