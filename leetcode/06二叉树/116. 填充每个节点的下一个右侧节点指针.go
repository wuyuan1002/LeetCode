package main

// 116. 填充每个节点的下一个右侧节点指针

// 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
//
// struct Node {
//  int val;
//  Node *left;
//  Node *right;
//  Node *next;
// }
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
//
// 初始状态下，所有 next 指针都被设置为 NULL。

// connect .
// 同 leetcode 102. 二叉树的层序遍历
// 层序遍历二叉树，遍历每一层的节点时将其next指针指向该层的下一个节点
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root} // 使用一个单端队列存储遍历过程中的左右节点 -- 初始化先将root入队

	for len(queue) > 0 { // 若当前层有节点，则遍历该层
		for count := len(queue); count > 0; count-- { // 开始遍历当前层，count：当前层节点个数
			// 获取当前层下一个节点并出队
			node := queue[0]
			queue = queue[1:]

			// 填充节点的next指针 -- 若为当前层最后一个节点则为nil，否则为当前层的下一个节点
			if count == 1 {
				node.Next = nil
			} else {
				node.Next = queue[0]
			}

			// 将当前节点的左右节点入队列 -- 为下一层的节点
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}

// Node .
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
