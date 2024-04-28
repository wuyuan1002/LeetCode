package main

// 剑指 Offer 25. 合并两个排序的链表

// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

// func main() {

// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var mergeHead *ListNode = nil
	if l1.Val >= l2.Val {
		mergeHead = l2
		mergeHead.Next = mergeTwoLists(l1, l2.Next)
	} else {
		mergeHead = l1
		mergeHead.Next = mergeTwoLists(l1.Next, l2)
	}
	return mergeHead
}

// 第二次做
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var list *ListNode = nil
	if l1.Val > l2.Val {
		list = l2
		list.Next = mergeTwoLists1(l1, l2.Next)
	} else {
		list = l1
		list.Next = mergeTwoLists1(l1.Next, l2)
	}
	return list
}
