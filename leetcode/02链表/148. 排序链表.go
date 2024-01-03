package main

// 148. 排序链表

// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

// sortList .
// 归并排序 O(nlogn) -- 自顶向下，或自底向上
// 不断找链表的中间节点，然后从中间将链表断开，将两端分别进行排序，然后再将两个排序的链表合并在一起
func sortList(head *ListNode) *ListNode {
	// 若传入的链表为空或只有一个节点，则直接返回head即可
	if head == nil || head.Next == nil {
		return head
	}

	// 使用快慢指针寻找链表中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next

	// 将链表从中点断开
	slow.Next = nil

	// 分别排序两个链表
	left := sortList(head)
	right := sortList(mid)

	// 合并两个排好序的链表
	return mergeTwoLists(left, right)
}
