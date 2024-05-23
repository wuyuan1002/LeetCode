package main

// LCR 123. 图书整理 I

// 书店店员有一张链表形式的书单，每个节点代表一本书，节点中的值表示书的编号。
// 为更方便整理书架，店员需要将书单倒过来排列，就可以从最后一本书开始整理，逐一将书放回到书架上。
// 请倒序返回这个书单链表。

// reverseBookList .
// Offer 06. 从尾到头打印链表
//
// 1. 先将链表元素挨个放入数组中，再将数组元素反转
// 2. 使用递归始终先遍历节点的后一个元素，把节点存入数组中
func reverseBookList(head *ListNode) []int {
	result := make([]int, 0)
	digui(head, &result)
	return result
}

// digui 始终先遍历当前节点的下一个节点，然后将当前节点存入数组中
func digui(head *ListNode, result *[]int) {
	if head == nil {
		return
	}
	digui(head.Next, result)
	*result = append(*result, head.Val)
}
