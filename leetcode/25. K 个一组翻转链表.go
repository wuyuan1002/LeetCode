package main

// 25. K 个一组翻转链表

func main() {

}

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
