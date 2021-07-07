package main

import (
	"fmt"
)

// 23. 合并K个升序链表

// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。

func main() {
	head1 := &ListNode2{
		Val: 1,
		Next: &ListNode2{
			Val: 4,
			Next: &ListNode2{
				Val:  5,
				Next: nil,
			},
		},
	}
	head2 := &ListNode2{
		Val: 1,
		Next: &ListNode2{
			Val: 3,
			Next: &ListNode2{
				Val:  4,
				Next: nil,
			},
		},
	}
	head3 := &ListNode2{
		Val: 2,
		Next: &ListNode2{
			Val:  6,
			Next: nil,
		},
	}

	aa := mergeKLists([]*ListNode2{head1, head2, head3})
	fmt.Println(aa)
}

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

// 类似offer51
// 1. 分治法 -- 类似归并排序，先解决前半部分链表合并，后解决后半部分链表合并，再合并两部分已合并好的链表；其实就是分割成两个链表合并问题
func mergeKLists(lists []*ListNode2) *ListNode2 {
	if lists == nil || len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	// 合并前半部分链表
	leftList := mergeKLists(lists[:len(lists)/2])
	// 合并后半部分链表
	rightList := mergeKLists(lists[len(lists)/2:])

	// 将合并好的两部分链表合并并返回
	return mergeTwoLists1(leftList, rightList)
}

func mergeTwoLists1(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var merge *ListNode2
	if l1.Val > l2.Val {
		merge = l2
		merge.Next = mergeTwoLists1(l1, l2.Next)
	} else {
		merge = l1
		merge.Next = mergeTwoLists1(l1.Next, l2)
	}

	return merge
}

// 2. 先合并数组中第1个和第2个链表，再将合并好的链表和第3个合并，再将合并好的链表和第4个合并...直到合并完数组中的所有链表
func mergeKLists1(lists []*ListNode2) *ListNode2 {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	// 合并好的链表
	merge := lists[0]
	for i := 1; i < len(lists); i++ {
		// 将合并好的链表和数组中下一个链表合并
		merge = mergeTwoLists1(merge, lists[i])
	}
	return merge
}
