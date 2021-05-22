package main

// 剑指 Offer 24. 反转链表

// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

func main() {

}

type ListNode3 struct {
	Val  int
	Next *ListNode3
}

func reverseList(head *ListNode3) *ListNode3 {
	if head == nil || head.Next == nil {
		return head
	}

	revers := head    // 已反转的节点
	node := head.Next // 当前节点，待反转
	next := node.Next // 当前节点的下一个节点，防止链表断裂

	head.Next = nil
	for node != nil {
		next = node.Next

		node.Next = revers
		revers = node
		node = next
	}
	return revers
}

// 第二次做
func reverseList1(head *ListNode3) *ListNode3 {
	if head == nil || head.Next == nil {
		return head
	}

	pre := head       // 前一个节点
	node := head.Next // 当前节点
	next := node.Next // 下一个节点

	head.Next = nil
	for node != nil {
		next = node.Next

		node.Next = pre

		pre = node
		node = next
	}

	return pre
}
