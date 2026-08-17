package main

// 24. 两两交换链表中的节点

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
// 你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

// swapPairs1 .
// 递归进行两两交换
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 设置新的头节点
	newHead := head.Next
	// 将旧头节点下一个节点的next设置为后续链表进行交换后的链表头
	head.Next = swapPairs1(newHead.Next)
	// 将新头节点的next指向旧的头节点 -- 交换旧的头节点head与head->next
	newHead.Next = head

	// 返回新的头节点
	return newHead
}

// swapPairs2
// leetcode 25. K个一组翻转链表
//
// K个一组翻转链表，当K==2时就是两个一组反转链表
func swapPairs2(head *ListNode) *ListNode {
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
	head.Next = swapPairs2(nextGroupHead)
	return newHead
}

// swapPairs3 .
// 设置虚拟头节点dummy，不断交换temp节点的后面两个节点，整个链表交换结束后返回最终的链表头节点
func swapPairs3(head *ListNode) *ListNode {
	// 设置虚拟头节点
	dummy := &ListNode{0, head}

	// 不断交换temp的后两个节点
	for temp := dummy; temp.Next != nil && temp.Next.Next != nil; {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}

	// 整个链表交换结束后，返回dummy指向的最终链表头节点
	return dummy.Next
}
