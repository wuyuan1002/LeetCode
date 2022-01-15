package main

// 148. 排序链表

// 给你链表的头结点head，请将其按 升序 排列并返回 排序后的链表 。
//
// 进阶：
// 你可以在O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

// func main() {

// }

// 类似Hot100 23, offer 51
// 1.插入排序 O(n2)
// 2.归并排序 O(nlogn) -- 自顶向下，或自底向上
func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	// 寻找链表中点
	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow

	return mergeTwoList(sort(head, mid), sort(mid, tail))
}

// 合并两个排序链表
func mergeTwoList(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	} else if head2 == nil {
		return head1
	}

	var merge *ListNode
	if head1.Val > head2.Val {
		merge = head2
		merge.Next = mergeTwoList(head1, head2.Next)
	} else {
		merge = head1
		merge.Next = mergeTwoList(head1.Next, head2)
	}

	return merge
}

type ListNode struct {
	Val  int
	Next *ListNode
}
