package main

// 206. 反转链表

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

func main() {

}

// 同offer 24
// 三指针
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 前一个节点、当前节点、下一个节点
	pre, node, next := head, head.Next, head.Next
	pre.Next = nil
	for node != nil {
		next = node.Next

		// 把当前节点的后继节点设置为前一个节点
		node.Next = pre

		pre = node
		node = next
	}

	return pre
}
