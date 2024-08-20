package main

// 83. 删除排序链表中的重复元素

// 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。

// deleteDuplicates83 .
// leetcode 82. 删除排序链表中的重复元素 II
//
// 本题要求将出现重复的数字保留一个在链表中，而82要求是将出现重复的数字一个不留全部删除
func deleteDuplicates83(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// 不断比较当前节点和前一个节点的值是否相等，若相等则将前一个节点的下一位指向
	// 当前节点的下一个节点（即将当前节点从链表中删除），若不相等则将前一个节点指向当前节点
	pre := head
	for node := head.Next; node != nil; node = node.Next {
		if node.Val == pre.Val {
			pre.Next = node.Next
		} else {
			pre = node
		}
	}

	return head
}
