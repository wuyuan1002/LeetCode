package main

// 117. 填充每个节点的下一个右侧节点指针 II

// 给定一个二叉树
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，
// 则将 next 指针设置为 NULL。
//
// 初始状态下，所有 next 指针都被设置为 NULL。
//
// 进阶：
// 你只能使用常量级额外空间。
// 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

func main() {

}

// 类似 116
// 不用队列存储每一层的节点来遍历那一层，而实使用那一层得next指针构成得链表遍历那一层
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	start := root // 一层的开始节点
	for start != nil {
		var nextStart, pre *Node // 下一层的开始节点、遍历的前一个节点
		for node := start; node != nil; node = node.Next {
			if node.Left != nil {
				if pre != nil {
					pre.Next = node.Left
				}
				pre = node.Left
				if nextStart == nil {
					nextStart = node.Left
				}
			}
			if node.Right != nil {
				if pre != nil {
					pre.Next = node.Right
				}
				pre = node.Right
				if nextStart == nil {
					nextStart = node.Right
				}
			}
		}
		start = nextStart
	}

	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
