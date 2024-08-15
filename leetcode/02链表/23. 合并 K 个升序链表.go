package main

// 23. 合并 K 个升序链表

// 给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。

// mergeKLists1 .
// 1. leetcode 21. 合并两个有序链表
// 先合并数组中第1个和第2个链表，再将合并好的链表和第3个合并，再将合并好的链表和第4个合并...直到合并完数组中的所有链表
func mergeKLists1(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	mergeList := lists[0]
	for i := 1; i < len(lists); i++ {
		mergeList = mergeTwoLists(mergeList, lists[i])
	}
	return mergeList
}

// mergeKLists2 .
// 2. 分治法
// 类似归并排序，先解决前半部分链表合并，后解决后半部分链表合并，再合并两部分已合并好的链表，
// 其实就是将所有链表不断分割两两合并，然后再整体合并
func mergeKLists2(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	return mergeMergeKLists(lists, 0, len(lists)-1)
}

// mergeMergeKLists 将两指针之间的所有链表不断分割，不断进行两个链表合并操作，最终将所有链表合并
func mergeMergeKLists(lists []*ListNode, l, r int) *ListNode {
	if l >= r {
		return lists[l]
	}

	mid := l + (r-l)/2
	list1 := mergeMergeKLists(lists, l, mid)
	list2 := mergeMergeKLists(lists, mid+1, r)

	return mergeTwoLists(list1, list2)
}
