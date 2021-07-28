package main

// 92. 反转链表 II

// 给你单链表的头指针 head 和两个整数left 和 right ，
// 其中left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回反转后的链表 。

func main() {

}

// 1. 遍历一次寻找到要翻转的子链表前后节点，翻转子链表还需要遍历一次子链表 -- 总共需要遍历两次链表
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left > right || left < 0 {
		return nil
	} else if left == right {
		return head
	}

	// 设置一个虚拟头节点
	var dummyNode = &ListNode{}
	dummyNode.Next = head

	// 寻找左右对应节点的前一个和后一个节点
	sub := right - left + 1
	l, r := dummyNode, head
	for i := 0; i < right; i++ {
		r = r.Next
		if i >= sub {
			l = l.Next
		}
	}

	n := reverselr(l.Next, r)
	l.Next.Next = r
	l.Next = n

	return dummyNode.Next
}

func reverselr(left, right *ListNode) *ListNode {
	var pre, node, next *ListNode = nil, left, left
	for node != right {
		next = node.Next

		node.Next = pre

		pre = node
		node = next
	}
	return pre
}

// 2. 寻找到左节点后继续向后遍历寻找右节点，过程中完成反转操作 -- 总共需要遍历一次
func reverseBetween1(head *ListNode, left, right int) *ListNode {

	// 设置dummyNode是这一类问题的一般做法
	dummyNode := &ListNode{}
	dummyNode.Next = head

	// 先找到左节点
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 向后遍历并翻转左右节点之间的节点
	node, next := pre.Next, pre.Next
	for i := 0; i < right-left; i++ {
		next = node.Next
		node.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
