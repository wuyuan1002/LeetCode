package main

// 86. 分隔链表

// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，
// 使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
// 你应当 保留 两个分区中每个节点的初始相对位置。

// partition .
// 创建两个空链表，遍历原链表，将节点按照与x的大小分别添加到新的两个链表中，最后将两个链表连接起来
func partition(head *ListNode, x int) *ListNode {
	// 新建两个链表分别存小于x的节点和大于等于x的节点
	smallHead, bigHead := &ListNode{}, &ListNode{}

	// 不断遍历链表节点，按照大小分别添加到两个链表中
	small, big := smallHead, bigHead
	for node := head; node != nil; node = node.Next {
		if node.Val < x {
			small.Next = node
			small = small.Next
		} else {
			big.Next = node
			big = big.Next
		}
	}

	// 将两个链表首尾相连 -- 大的连到小的后面
	small.Next = bigHead.Next
	big.Next = nil

	return smallHead.Next
}
