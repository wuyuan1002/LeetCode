package main

// 206. 反转链表

func main() {

}

// 同offer 24
// 三指针
func reverseList(head *ListNode) *ListNode {
	pre, node, next := (*ListNode)(nil), head, head
	for node != nil {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}
