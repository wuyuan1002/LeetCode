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

	return sort1(head, nil)
}

func sort1(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow
	return mergeTwoList(sort1(head, mid), sort1(mid, tail))
}

// 合并两个排序链表
func mergeTwoList(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	} else if head2 == nil {
		return head1
	}

	merge := &ListNode{}
	if head1.Val < head2.Val {
		merge.Val = head1.Val
		merge.Next = mergeTwoList(head1.Next, head2)
	} else {
		merge.Val = head2.Val
		merge.Next = mergeTwoList(head1, head2.Next)
	}

	return merge
}
