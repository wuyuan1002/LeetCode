package main

// 142. 环形链表 II

func main() {

}

// 1. 遍历所有节点，使用map记录已经访问过的节点
// 2. 快慢指针 -- 两指针相遇后，再定义一个慢指针从链表头出发，当快慢指针再次相遇时，它指向的就是环的入口
func detectCycle(head *ListNode) *ListNode {
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
			return slow
		}
	}
	return nil
}
