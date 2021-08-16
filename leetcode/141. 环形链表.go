package main

// 141. 环形链表

func main() {

}

// 1. 遍历所有节点，使用map记录已经访问过的节点
// 2. 快慢指针
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}
