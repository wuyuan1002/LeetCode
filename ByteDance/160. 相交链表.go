package main

// 160. 相交链表

// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。
// 如果两个链表没有交点，返回 null 。

func main() {

}

// 1.把两个链表分别入栈，然后同时从栈中弹出元素，直到第一个不相等的元素
// 2.将第一个链表全部入map，之后从头到尾遍历第二个链表，判断map中是否存在节点
// 3.双指针
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}

		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}

	return p
}
