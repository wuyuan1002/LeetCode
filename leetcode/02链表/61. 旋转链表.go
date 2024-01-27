package main

// 61. 旋转链表

// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

// rotateRight .
// 先把链表首位相连，之后根据k计算断开的位置
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}

	// 首位相连链表同时计算链表长度
	count := 0
	for node := head; node != nil; node = node.Next {
		count++
		if node.Next == nil {
			node.Next = head
			break
		}
	}

	// 计算应该从哪里断开
	cutIndex := count - k%count

	// 在指定位置断开链表
	cutNode := head
	for i := 0; i < cutIndex-1; i++ { // 先找到要切断位置的前一个节点
		cutNode = cutNode.Next
	}
	newHead := cutNode.Next
	cutNode.Next = nil

	return newHead
}
