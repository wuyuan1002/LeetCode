package main

// 148. 排序链表

// 给你链表的头结点head，请将其按 升序 排列并返回 排序后的链表 。
//
// 进阶：
// 你可以在O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

func main() {

}

// 类似Hot100 23, offer 51
// 1.插入排序 O(n2)
// 2.归并排序 O(nlogn) -- 自顶向下，或自底向上

// 不是自己写的, 是leetcode的答案
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast.Next != tail && fast.Next.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow
	return mergeTwoList(sort(head, mid), sort(mid, tail))
}

func mergeTwoList(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}
