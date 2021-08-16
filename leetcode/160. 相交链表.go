package main

// 160. 相交链表

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
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}

		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}

	return p
}
