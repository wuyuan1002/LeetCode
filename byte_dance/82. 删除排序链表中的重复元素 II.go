package main

// 82. 删除排序链表中的重复元素 II

// 存在一个按升序排列的链表，给你这个链表的头节点 head ，
// 请你删除链表中所有存在数字重复情况的节点，只保留原始链表中没有重复出现的数字。
//
// 返回同样按升序排列的结果链表。

// func main() {

// }

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head} // 虚拟头节点
	node := dummy
	for node.Next != nil && node.Next.Next != nil {
		if node.Next.Val == node.Next.Next.Val {
			val := node.Next.Val
			for node.Next != nil && node.Next.Val == val {
				node.Next = node.Next.Next
			}
		} else {
			node = node.Next
		}
	}

	return dummy.Next
}
