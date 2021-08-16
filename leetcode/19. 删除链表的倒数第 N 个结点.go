package main

// 19. 删除链表的倒数第 N 个结点

func main() {

}

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
