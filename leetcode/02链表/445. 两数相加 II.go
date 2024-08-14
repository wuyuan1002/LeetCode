package main

// 445. 两数相加 II

// 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。
// 它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。

// addTwoNumbers445 .
// leetcode 2.两数相加
//
// 1. 反转两链表，然后顺序遍历链表进行各位相加
// 2. 将两个链表数字分别存入两个栈中，然后不断将两个栈中的数字出栈进行各位相加
func addTwoNumbers445(l1 *ListNode, l2 *ListNode) *ListNode {

	// 指向总结果的头节点
	result := (*ListNode)(nil)

	// 将两个链表数字分别存入两个栈中
	stack1, stack2 := make([]int, 0), make([]int, 0)
	for ; l1 != nil; l1 = l1.Next {
		stack1 = append(stack1, l1.Val)
	}
	for ; l2 != nil; l2 = l2.Next {
		stack2 = append(stack2, l2.Val)
	}

	// 不断将两个栈中的数字出栈进行各位相加
	carry := 0 // 进位
	for len(stack1) > 0 || len(stack2) > 0 || carry > 0 {
		// 将数字出栈并计算当前位的和
		sum := carry
		if len(stack1) > 0 {
			sum += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			sum += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}

		// 将当前得数插入到结果链表头
		result = &ListNode{Val: sum % 10, Next: result}

		// 计算当前位的进位
		carry = sum / 10
	}

	return result
}
