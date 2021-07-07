package main

// 2. 两数相加

// 给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0开头。

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := &ListNode{}
	res := head

	carry := 0       // 进位
	num := 0         // 某一位得数
	n1, n2 := l1, l2 // 正在相加的数
	for n1 != nil || n2 != nil {
		sum := carry
		if n1 != nil {
			sum += n1.Val
		}
		if n2 != nil {
			sum += n2.Val
		}

		carry = sum / 10
		num = sum % 10

		node := &ListNode{
			Val: num,
		}
		res.Next = node
		res = res.Next

		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
	}

	if carry > 0 {
		node := &ListNode{
			Val:  carry,
			Next: nil,
		}

		res.Next = node
	}

	return head.Next
}
