package main

// 19. 删除链表的倒数第 N 个结点

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 进阶：你能尝试使用一趟扫描实现吗？

// func main() {

// 	head := &ListNode1{
// 		Val:  5,
// 		Next: nil,
// 	}
// 	fmt.Println(removeNthFromEnd(head, 1).Val)
// }

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

// 同Offer 22
// 定义两个指针，一个指针先走k-1步，之后两个指针同时向后走，当前面指针为nil时，后一个指针就正好是倒数第k个
func removeNthFromEnd(head *ListNode1, n int) *ListNode1 {
	if head == nil || n <= 0 {
		return head
	}

	l, r := head, head

	// 先让前面的指针向前走n步
	for i := 0; i < n; i++ {
		if r == nil {
			return head
		}
		r = r.Next
	}

	// 两个指针同时向前移动，直到前面的指针r的next为nil
	for r != nil && r.Next != nil {
		l = l.Next
		r = r.Next
	}

	if r == nil {
		// 倒数第n个节点恰巧是头节点
		return head.Next
	}

	// 后面指针指向的下一个节点就是倒数第n个节点 -- 删除倒数第n个节点
	l.Next = l.Next.Next

	return head
}
