package main

// 剑指 Offer 32 - I. 从上到下打印二叉树

// 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

// func main() {

// }

func levelOrder(root *TreeNode) []int {
	result := make([]int, 0)
	if root != nil {
		queue := []*TreeNode{root}
		var node *TreeNode
		for i := 0; i < len(queue); i++ {
			node = queue[i]
			result = append(result, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}

// 第二次做 -- 广度优先遍历树 -- 都需要一个队列作为辅助空间
func levelOrder4(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	queue := []*TreeNode{root}

	var node *TreeNode
	for i := 0; i < len(queue); i++ {
		node = queue[i]
		result = append(result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}
