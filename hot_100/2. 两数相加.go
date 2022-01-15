package main

// 2.两数相加

// 给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，
// 并且每个节点只能存储一位数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0开头。

// func main() {
// 	headA := &ListNode{
// 		Val: 9,
// 		Next: &ListNode{
// 			Val: 9,
// 			Next: &ListNode{
// 				Val:  9,
// 				Next: nil,
// 			},
// 		},
// 	}

// 	headB := &ListNode{
// 		Val: 9,
// 		Next: &ListNode{
// 			Val:  9,
// 			Next: nil,
// 		},
// 	}
// 	node1 := addTwoNumbers(headA, headB)
// 	fmt.Println(node1)
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

// 由于两个链表是是逆序存储数字的，秩序从头遍历两个链表相加即可
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := &ListNode{} // 结果
	res := head         // 当前结果节点

	carry := 0             // 进位
	num := 0               // 相加后当前位数值
	node1, node2 := l1, l2 // l1和l2上当前进行相加的两个数字
	for node1 != nil || node2 != nil {
		// 计算两数之和
		sum := 0
		if node1 != nil {
			sum += node1.Val
		}
		if node2 != nil {
			sum += node2.Val
		}
		sum += carry

		// 计算进位和当前位数值
		carry = sum / 10
		num = sum % 10

		// 构造节点
		node := &ListNode{
			Val: num,
		}

		// 将结果连接到结果链表中
		res.Next = node
		res = res.Next

		if node1 != nil {
			node1 = node1.Next
		}
		if node2 != nil {
			node2 = node2.Next
		}
	}

	// 处理最后一个进位
	if carry > 0 {
		node := &ListNode{
			Val: carry,
		}

		// 将结果连接到结果链表中
		res.Next = node
		res = res.Next
	}

	return head.Next
}
