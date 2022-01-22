package main

// 24. 两两交换链表中的节点

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// func main() {

// }

// 类似 25 K个一组翻转链表, 本题等于两个一组翻转链表
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nextStartNode := head
	for i := 0; i < 2; i++ {
		if nextStartNode == nil {
			return head
		}
		nextStartNode = nextStartNode.Next
	}

	newHead := reverse1(head, nextStartNode)

	head.Next = swapPairs(nextStartNode)

	return newHead
}

func reverse1(head, nextStartNode *ListNode) *ListNode {
	var pre, node, next *ListNode = nil, head, head
	for next != nextStartNode {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}
