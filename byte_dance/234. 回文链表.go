package main

// 234. 回文链表

// 请判断一个链表是否为回文链表。

func main() {

}

// 1.第一次遍历链表，将每个节点放入栈中，再一次遍历链表，分别与栈顶元素比较
// 2.将链表遍历放入数组中，之后使用双指针从头尾遍历比较
// 3.先使用快慢指针找到链表中间的节点，之后将前半部分或后半部分翻转，之后使用双指针遍历前后两个部分进行比较
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 快慢指针寻找中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 将后半部分反转
	end := reversList(slow.Next)

	// 从两边遍历链表，分别比较
	res := true
	p1, p2 := head, end
	for res && p2 != nil {
		if p1.Val != p2.Val {
			res = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 将后半部分反转回去
	slow.Next = reversList(end)

	return res
}

// 反转链表
func reversList(head *ListNode) *ListNode {
	var pre, node, next *ListNode = nil, head, head
	for node != nil {
		next = node.Next

		node.Next = pre

		pre = node
		node = next
	}
	return pre
}
