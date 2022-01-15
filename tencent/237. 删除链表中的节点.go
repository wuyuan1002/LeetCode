package main

// 237. 删除链表中的节点

// 请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。

// func main() {

// }

// 类似Offer 18. 删除链表的节点
func deleteNode(node *ListNode) {
	if node == nil || node.Next == nil {
		return
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
