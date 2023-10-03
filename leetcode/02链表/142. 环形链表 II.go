package leetcode

// 142. 环形链表 II

// 给定一个链表，返回链表开始入环的第一个节点。如果链表无环，则返回null。
//
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
// 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
//
// 说明：不允许修改给定的链表。
//
// 进阶：
// 你是否可以使用 O(1) 空间解决此题？

// detectCycle .
// 同 141. 环形链表
// 使用快慢指针确认链表有环后, 两指针一个从head开始一个从相遇位置开始, 继续向前移动, 直到相遇时相遇点就是入环的起点
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}
