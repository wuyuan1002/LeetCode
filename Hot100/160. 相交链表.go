package main

// 给你两个单链表的头节点headA 和 headB ，
// 请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
//
// 题目数据 保证 整个链式结构中不存在环。
// 注意，函数返回结果后，链表必须 保持其原始结构 。

func main() {

}

// 同offer 52
// 1.把两个链表分别入栈，然后同时弹出元素，直到第一个不相等的元素
// 2.将第一个链表全部入map，之后从头到尾遍历第二个链表，判断map中是否存在节点
// 3.双指针
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	i, j := headA, headB
	for i != j {
		if i == nil {
			i = headB
		} else {
			i = i.Next
		}

		if j == nil {
			j = headA
		} else {
			j = j.Next
		}
	}

	return i
}
