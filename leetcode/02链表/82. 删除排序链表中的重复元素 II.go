package main

// 82. 删除排序链表中的重复元素 II

// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。

// deleteDuplicates82 .
// leetcode 83. 删除排序链表中的重复元素
//
// 本题要求将出现重复的数字一个不留全部删除，而83要求是将出现重复的数字保留一个在链表中
func deleteDuplicates82(head *ListNode) *ListNode {
	// 虚拟头节点
	dummy := &ListNode{Next: head}

	// node初始为虚拟头节点，不断判断node的后两个节点值是否为重复数字，若不重复则将node向后移动一位然后继续判断其后两个值是否为重复数字，
	// 若node的后两个节点为重复数字，则需要将所有重复的数字从链表中移除，也就是将node的下一位指向移除当前重复数字后的首个节点
	node := dummy
	for node.Next != nil && node.Next.Next != nil {
		if node.Next.Val == node.Next.Next.Val {
			// node的后两个节点为重复数字，需要将所有重复的数字移除 -- 不断判断node.next是否还是重复数字并向后移动node.next
			val := node.Next.Val
			for node.Next != nil && node.Next.Val == val {
				node.Next = node.Next.Next
			}
		} else {
			// node的后两个节点不是重复数字，将node向后移动一位，继续判断下一位的后两个数字是否重复
			node = node.Next
		}
	}

	return dummy.Next
}
