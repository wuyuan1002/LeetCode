package main

// 237. 删除链表中的节点

// 请编写一个函数，用于 删除单链表中某个特定节点 。在设计函数时需要注意，
// 你无法访问链表的头节点head ，只能直接访问 要被删除的节点 。
//
// 题目数据保证需要删除的节点 不是末尾节点 。

func main() {

}

// 将当前节点的下一个节点的值拷贝到当前节点，之后删掉当前节点的下一个节点
func deleteNode(node *ListNode) {
	if node == nil || node.Next == nil {
		return
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
