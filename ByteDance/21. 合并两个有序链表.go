package main

// 21. 合并两个有序链表

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

func main() {

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var merge *ListNode
	if l1.Val > l2.Val {
		merge = l2
		merge.Next = mergeTwoLists(l1, l2.Next)
	} else {
		merge = l1
		merge.Next = mergeTwoLists(l1.Next, l2)
	}

	return merge
}
