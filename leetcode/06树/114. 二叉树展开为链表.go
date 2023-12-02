package main

// 114. 二叉树展开为链表

// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
//
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。

// flatten .
// 右左根后序遍历二叉树 -- 因为结果要与前序遍历一致，而递归相当于栈
func flatten(root *TreeNode) {

	// 指向前一个遍历的节点
	pre := (*TreeNode)(nil)

	// 后序遍历二叉树 -- 因为结果要与前序遍历一致，而递归相当于栈
	var dfsFlatten func(node *TreeNode)
	dfsFlatten = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfsFlatten(node.Right)
		dfsFlatten(node.Left)

		node.Left = nil
		node.Right = pre
		pre = node
	}

	// 开始遍历跟节点
	dfsFlatten(root)
}
