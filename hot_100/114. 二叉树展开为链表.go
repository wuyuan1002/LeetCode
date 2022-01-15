package main

// 114. 二叉树展开为链表

// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，
// 而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。

// func main() {

// }

// 类似 offer 36

var preNode *TreeNode = nil // 遍历到的前一个节点

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 后序遍历二叉树 -- 因为结果要与前序遍历一致，而递归相当于栈
	flatten(root.Right)
	flatten(root.Left)

	// 将当前节点的左节点置为空
	root.Left = nil

	// 右节点指向上一个遍历到的节点
	root.Right = preNode
	// 更新上一个遍历到的节点
	preNode = root
}
