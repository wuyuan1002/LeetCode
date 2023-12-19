package main

// 24. 两两交换链表中的节点

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
// 你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

// swapPairs .
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 设置新的头节点
	newHead := head.Next
	// 将头节点下一个节点的next设置为后续链表进行交换后的链表头
	head.Next = swapPairs(newHead.Next)
	// 将当前头节点的next指向旧的头节点（交换旧的头节点head与head->next）
	newHead.Next = head
	// 返回新的头节点
	return newHead
}

// swapPairs1
// 同 leetcode 25. K个一组翻转链表 -- 当K==2时就是两个一组反转链表
func swapPairs1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	nextGroupHead := head
	for i := 0; i < 2; i++ {
		if nextGroupHead == nil {
			return head
		}
		nextGroupHead = nextGroupHead.Next
	}

	newHead := reverse(head, nextGroupHead)
	head.Next = swapPairs1(nextGroupHead)
	return newHead
}
