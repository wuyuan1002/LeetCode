package main

// 23. 合并K个升序链表

// 给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。

func main() {

}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	merge1 := mergeKLists(lists[:len(lists)/2])
	merge2 := mergeKLists(lists[len(lists)/2:])

	return mergeList(merge1, merge2)
}

// 合并两个链表
func mergeList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var merge *ListNode
	if l1.Val > l2.Val {
		merge = l2
		merge.Next = mergeList(l1, l2.Next)
	} else {
		merge = l1
		merge.Next = mergeList(l1.Next, l2)
	}

	return merge
}
