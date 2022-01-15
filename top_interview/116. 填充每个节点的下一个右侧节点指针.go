package main

// 116. 填充每个节点的下一个右侧节点指针

// 给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
//
// struct Node {
//  int val;
//  Node *left;
//  Node *right;
//  Node *next;
// }
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，
// 则将 next 指针设置为 NULL。
//
// 初始状态下，所有next 指针都被设置为 NULL。

// func main() {
// 	root := &Node{
// 		Val: 1,
// 		Left: &Node{
// 			Val: 2,
// 			Left: &Node{
// 				Val: 4,
// 			},
// 			Right: &Node{
// 				Val: 5,
// 			},
// 			Next: nil,
// 		},
// 		Right: &Node{
// 			Val: 3,
// 			Left: &Node{
// 				Val: 6,
// 			},
// 			Right: &Node{
// 				Val: 7,
// 			},
// 			Next: nil,
// 		},
// 		Next: nil,
// 	}

// 	connect(root)
// }

// 类似 117
// 二叉树的层序遍历
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}
	num := 1     // 当前层剩余节点数
	nextNum := 0 // 下一层节点数
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		num--

		if node.Left != nil {
			queue = append(queue, node.Left)
			nextNum++
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nextNum++
		}

		if num == 0 {
			num = nextNum
			nextNum = 0

			// 一层的最后一个节点，next置为nil
			node.Next = nil
		} else {
			// 每一层节点的next置为当前层的下一个节点
			node.Next = queue[i+1]
		}
	}

	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
