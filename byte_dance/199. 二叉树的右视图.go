package main

// 199. 二叉树的右视图

// 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

// func main() {

// }

// 层级遍历二叉树，把每一层的最后一个入结果集
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)

	queue := []*TreeNode{root}
	count := 1     // 当前层的个数
	nextCount := 0 // 下一层的个数
	var node *TreeNode
	for i := 0; i < len(queue); i++ {
		node = queue[i]
		count--

		if node.Left != nil {
			queue = append(queue, node.Left)
			nextCount++
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nextCount++
		}

		// 一行的最后一个元素
		if count == 0 {
			result = append(result, node.Val)
			count = nextCount
			nextCount = 0
		}
	}

	return result
}
