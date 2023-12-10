package main

// 2.两数相加

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，
// 并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// addTwoNumbers .
// 由于两个链表是是逆序存储数字的，只需从头遍历两个链表相加即可
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // 总结果虚拟头节点
	preNode := dummy     // 上一位的结果节点

	node1, node2 := l1, l2 // 正在参与计算的两个节点
	carry := 0             // 进位

	for node1 != nil || node2 != nil || carry != 0 {
		// 计算两数之和 -- sum = num1 + num2 + carry
		sum := carry
		if node1 != nil {
			sum += node1.Val
		}
		if node2 != nil {
			sum += node2.Val
		}

		// 计算当前位得数并构造当前位节点、计算进位
		curNode := &ListNode{Val: sum % 10}
		carry = sum / 10

		// 将当前位结果加入到结果链表中
		preNode.Next = curNode
		preNode = preNode.Next

		// 向后移动参与计算的数字
		if node1 != nil {
			node1 = node1.Next
		}
		if node2 != nil {
			node2 = node2.Next
		}
	}

	// 返回结果头节点
	return dummy.Next
}
