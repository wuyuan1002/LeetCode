package main

// 147. 对链表进行插入排序

// 给定单个链表的头 head ，使用 插入排序 对链表进行排序，并返回 排序后链表的头 。
//
// 插入排序 算法的步骤:
// 插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
// 每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
// 重复直到所有输入数据插入完为止。

// insertionSortList .
// 使用一个变量记录已排好序的尾节点，不断遍历每一个节点将其插入到正确的位置
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 虚拟头节点
	dummy := &ListNode{Next: head}

	// 已排好序的尾节点
	sortedTail := head

	// 遍历每一个节点，将其插入到正确的位置
	for current := head.Next; current != nil; current = sortedTail.Next {
		// 若要插入的节点正好应该在整个排序链表的最后，则直接更新尾节点
		if current.Val >= sortedTail.Val {
			sortedTail = current
			continue
		}

		// 遍历排好序的链表，寻找要被插入节点的前一个节点pre
		pre := dummy
		for pre.Next.Val <= current.Val {
			pre = pre.Next
		}

		// 将当前节点插入到pre的后面
		sortedTail.Next = current.Next
		current.Next = pre.Next
		pre.Next = current
	}

	return dummy.Next
}
