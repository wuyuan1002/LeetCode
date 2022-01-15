package main

import "fmt"

// 非递归遍历二叉树

// func main() {

// }

// 前序遍历
func preOrder(root *TreeNode) {
	// 使用栈保存已被访问的根节点
	stack := make([]*TreeNode, 0)
	node := root
	for node != nil || len(stack) > 0 {
		if node != nil {
			// 若当前节点不为空

			// 访问到当前节点并将当前节点入栈
			fmt.Println(node.Val)
			stack = append(stack, node)

			// 访问当前节点的左节点
			node = node.Left
		} else {
			// 若当前节点为空，说明已经访问到了最左面

			// 取栈中最近被访问的节点的右子树
			node = stack[len(stack)-1].Right
			// 将最近已被访问的根节点出栈
			stack = stack[:len(stack)-1]
		}
	}
}

// 中序遍历
func inOrder(root *TreeNode) {
	// 使用栈保存需要被访问的根节点
	stack := make([]*TreeNode, 0)
	node := root
	for node != nil || len(stack) > 0 {
		if node != nil {
			// 若当前节点不为空，先一股脑将左子树都入栈

			// 先不访问当前节点，将当前节点入栈
			stack = append(stack, node)

			// 访问当前节点的左节点
			node = node.Left
		} else {
			// 若当前节点为空，说明已经访问到了最左面

			// 取栈中最近未被访问的根节点
			node = stack[len(stack)-1]
			// 将最近未被访问的根节点出栈
			stack = stack[:len(stack)-1]

			// 访问到根节点
			fmt.Println(node.Val)

			// 访问当前节点的右子树
			node = node.Right
		}
	}
}

// 后序遍历
// 1. 传统非递归遍历方法
// 2. 使用根右左的前序遍历方法遍历，然后在反转数组即可
func laterOrder(root *TreeNode) {
	// 使用栈保存需要被访问的根节点
	stack := make([]*TreeNode, 0)
	node, pre := root, (*TreeNode)(nil) // 当前节点、前一个被访问的节点
	for node != nil || len(stack) > 0 {
		if node != nil {
			// 若当前节点不为空，先一股脑将左子树都入栈

			// 先不访问当前节点，将当前节点入栈
			stack = append(stack, node)

			// 访问当前节点的左节点
			node = node.Left
		} else {
			// 若当前节点为空，说明已经访问到了最左面

			// 取栈中最近未被访问的根节点
			node = stack[len(stack)-1]
			if node.Right == nil || node.Right == pre {
				// 若栈顶根节点的右子树为空或刚被访问过，则访问当前根节点

				// 将最近未被访问的根节点出栈
				stack = stack[:len(stack)-1]

				fmt.Println(node.Val)

				pre = node
				node = nil
			} else {
				// 当前节点的右节点还未被访问，先去访问它的右子树
				node = node.Right
			}
		}
	}
}
