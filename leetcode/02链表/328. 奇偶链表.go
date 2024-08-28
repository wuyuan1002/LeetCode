package main

// 328. 奇偶链表

// 给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
// 第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
//
// 请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
// 你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。

// oddEvenList .
// 一次遍历链表，使用两个指针分别指向当前遍历的偶数节点和奇数节点，，
// 遍历时将偶数节点和奇数节点分别串联成单独的链表，遍历结束后再将偶数链表添加到奇数链表尾部
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// 当前遍历到的奇数节点和偶数节点
	odd, even := head, head.Next
	// 偶数链表的链表头节点
	evenHead := even

	// 不断遍历链表节点，将偶数节点和奇数节点分别串联成单独的链表
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	// 将偶数链表添加到奇数链表的尾部
	odd.Next = evenHead

	return head
}
