package main

// 25. K 个一组翻转链表

// 给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
// k是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 进阶：
// 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

// func main() {

// }

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return head
	}

	nextGroupStartNode := head
	for i := 0; i < k; i++ {
		if nextGroupStartNode == nil {
			return head
		}
		nextGroupStartNode = nextGroupStartNode.Next
	}

	newHead := reverse(head, nextGroupStartNode)
	head.Next = reverseKGroup(nextGroupStartNode, k)
	return newHead
}

func reverse(head *ListNode, nextGroupStartNode *ListNode) *ListNode {
	pre, node, next := (*ListNode)(nil), head, head
	for node != nextGroupStartNode {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}

// 第二次做
func reverseKGroup1(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return head
	}

	nextGroupStartNode := head
	for i := 0; i < k; i++ {
		if nextGroupStartNode == nil {
			return head
		}
		nextGroupStartNode = nextGroupStartNode.Next
	}

	newHead := reverse11(head, nextGroupStartNode)
	head.Next = reverseKGroup1(nextGroupStartNode, k)
	return newHead
}

func reverse11(head *ListNode, nextGroupStartNode *ListNode) *ListNode {
	pre, node, next := (*ListNode)(nil), head, head
	for node != nextGroupStartNode {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}
