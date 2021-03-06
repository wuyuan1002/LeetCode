package main

// 148. 排序链表

// 给你链表的头结点head，请将其按 升序 排列并返回 排序后的链表 。
//
// 进阶：
// 你可以在O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

// func main() {

// }

// 类似Hot100 23, offer 51
// 1.插入排序 O(n2)
// 2.归并排序 O(nlogn) -- 自顶向下，或自底向上
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	return nil
}
