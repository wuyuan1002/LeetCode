package main

// 142. 环形链表 II

// 给定一个链表，返回链表开始入环的第一个节点。如果链表无环，则返回null。
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
// 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
//
// 说明：不允许修改给定的链表。
//
// 进阶：
// 你是否可以使用 O(1) 空间解决此题？
var (
	node1 = &ListNode3{
		Val:  3,
		Next: nil,
	}
	node2 = &ListNode3{
		Val:  2,
		Next: nil,
	}
	node3 = &ListNode3{
		Val:  0,
		Next: nil,
	}
	node4 = &ListNode3{
		Val:  4,
		Next: nil,
	}
)

// func main() {
// 	node1.Next = node2
// 	node2.Next = node3
// 	node3.Next = node4
// 	node4.Next = node2
// 	fmt.Println(detectCycle(node1).Val)
// }

// 1. 遍历所有节点，使用map记录已经访问过的节点
// 2. 快慢指针 -- 两指针相遇后，再定义一个慢指针从链表头出发，当快慢指针再次相遇时，它指向的就是环的入口
func detectCycle(head *ListNode3) *ListNode3 {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return fast
		}
	}
	return nil
}

type ListNode3 struct {
	Val  int
	Next *ListNode3
}
