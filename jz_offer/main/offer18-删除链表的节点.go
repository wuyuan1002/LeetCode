package main

// 剑指 Offer 18. 删除链表的节点
// 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
// 返回删除后的链表的头节点。

// func main() {
// }

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

// 顺序遍历链表，找到要删除的节点，把它删除 -- 时间复杂度 O(n)
func deleteNode(head *ListNode1, val int) *ListNode1 {
	i := head
	pre := head // 指向i的上一个
	for ; i != nil; i = i.Next {
		if i.Val == val { // 找到了要删除的节点
			pre.Next = i.Next
			if i == pre { // 若要删除的是头节点
				head = i.Next
			}
			i = nil // 删除当前节点
			break
		}
		pre = i
	}
	return head
}

// 第二次做 -- 直接把要删除的节点的下一个的值复制给要删除的节点，再把要删除节点的next指针指向下一个的下一个，这里需要保证要删除的节点必须在链表中，时间复杂度 O(1)
func deleteNode1(head *ListNode1, deleted *ListNode1) *ListNode1 {
	return nil
}
