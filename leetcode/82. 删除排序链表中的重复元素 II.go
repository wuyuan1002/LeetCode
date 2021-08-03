package main

// 82. 删除排序链表中的重复元素 II

func main() {

}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 虚拟头节点
	dummy := &ListNode{
		Next: head,
	}

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

type ListNode struct {
	Val  int
	Next *ListNode
}
