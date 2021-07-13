package main

// 83. 删除排序链表中的重复元素

// 存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素只出现一次 。
//
// 返回同样按升序排列的结果链表。

func main() {

}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := head
	for node := head.Next; node != nil; node = node.Next {
		if pre.Val == node.Val {
			pre.Next = node.Next
		} else {
			pre = node
		}
	}

	return head
}
