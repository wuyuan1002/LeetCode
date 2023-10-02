package main

// 19. 删除链表的倒数第 N 个结点

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 进阶：你能尝试使用一趟扫描实现吗？

// func main() {

// }

// 同Offer 22
// 定义两个指针，一个指针先走k-1步，之后两个指针同时向后走，当前面指针为nil时，后一个指针就正好是倒数第k个
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

// 第二次做
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	// 设置虚拟头节点
	dummy := &ListNode{Next: head}
	// 快慢指针
	l, r := dummy, head
	for i := 0; r != nil; i++ {
		r = r.Next
		if i < n {
			// 先让快指针向前走k步
			continue
		}
		l = l.Next
	}

	// 此时l的next就是倒数第k个节点 -- 删掉倒数第k个节点
	l.Next = l.Next.Next
	return dummy.Next
}
