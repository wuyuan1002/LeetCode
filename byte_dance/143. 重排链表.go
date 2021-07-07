package main

// 143. 重排链表

// 给定一个单链表L：L0→L1→…→Ln-1→Ln ，
// 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

func main() {

}

// 1. 将链表放入数组，这样就可以访问前一个元素了，然后进行合并
// 2. 找到链表中点，将后半段翻转，再和前半段做合并

// 1.
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	nodes := make([]*ListNode, 0)
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

// 2. 运行没过
func reorderList1(head *ListNode) {
	if head == nil {
		return
	}

	// 获取中间节点
	midNode := getMiddleNode(head)
	// 反转后半段链表
	tail := reverseList1(midNode)

	// 合并链表
	l1, l2 := head, tail
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

// 获取链表中间节点
func getMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 反转链表
func reverseList1(head *ListNode) *ListNode {
	var pre, node, next *ListNode = nil, head, head
	for node != nil {
		next = node.Next

		node.Next = pre

		pre = node
		node = next
	}
	return pre
}
