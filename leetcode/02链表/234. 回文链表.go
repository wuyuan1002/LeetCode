package main

// 234. 回文链表

// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

// isPalindrome .
// 1.第一次遍历链表，将每个节点放入栈中，再一次遍历链表，分别与栈顶元素比较
// 2.将链表遍历放入数组中，之后使用双指针从头尾遍历比较
// 3.先使用快慢指针找到链表中间的节点，之后将前半部分或后半部分翻转，之后使用双指针遍历前后两个部分进行比较
func isPalindrome(head *ListNode) bool {
	// 1. 快慢指针寻找链表中点 -- 偶数找到的是前一个
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. 反转链表后半段
	tail := reversList234(slow.Next)

	// 3. 挨个遍历未反转的前半段和反转后的后半段各个节点值是否相同
	result := true
	p, q := head, tail
	for result && q != nil {
		if p.Val != q.Val {
			result = false
		}
		p = p.Next
		q = q.Next
	}

	// 4. 恢复链表
	slow.Next = reversList234(tail)

	// 5. 返回结果
	return result
}

// reversList234 反转链表
func reversList234(head *ListNode) *ListNode {
	pre, node, next := (*ListNode)(nil), head, head
	for node != nil {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}
