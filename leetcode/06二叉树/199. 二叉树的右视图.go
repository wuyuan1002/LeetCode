package main

// 199. 二叉树的右视图

// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

// rightSideView .
// 同 leetcode 102. 二叉树的层序遍历
// 层序遍历二叉树，在遍历到每一层的最后一个元素时将节点入结果集
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)   // 结果集
	queue := []*TreeNode{root} // 使用一个单端队列存储遍历过程中的左右节点 -- 初始化先将root入队

	for len(queue) > 0 { // 若当前层有节点，则遍历该层
		for count := len(queue); count > 0; count-- { // 开始遍历当前层，count：当前层节点个数
			// 获取当前层下一个节点并出队
			node := queue[0]
			queue = queue[1:]

			// 若当前节点时当前层的最后一个节点则入结果集
			if count == 1 {
				result = append(result, node.Val)
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

	return result
}
