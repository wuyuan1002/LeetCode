package main

// 141. 环形链表

// 给定一个链表，判断链表中是否有环。
//
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
// 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
// 如果链表中存在环，则返回 true 。 否则，返回 false 。
//
// 进阶：
//
// 你能用 O(1)（即，常量）内存解决此问题吗？

// hasCycle .
// leetcode 142. 环形链表 II
//
// 快慢指针
// 分别定义fast和slow指针，从头结点出发，fast指针每次移动两个节点，
// slow指针每次移动一个节点，如果fast和slow指针在途中相遇，说明这个链表有环
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}
