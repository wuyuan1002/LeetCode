package main

// 513. 找树左下角的值

// 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//
// 假设二叉树中至少有一个节点。

// findBottomLeftValue .
// 同 leetcode 102. 二叉树的层序遍历
// 层级遍历二叉树，每次记录当层的第一个节点值，遍历完后，结果就为最后一层的第一个节点的值
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := -1               // 结果
	queue := []*TreeNode{root} // 使用一个单端队列存储遍历过程中的左右节点 -- 初始化先将root入队

	for len(queue) > 0 { // 若当前层有节点，则遍历该层
		// 获取当前层的节点个数
		count := len(queue)
		// 开始遍历当前层
		for i := 0; i < count; i++ {
			node := queue[0]
			queue = queue[1:]

			// 若当前节点为当前层的第一个节点，则更新总结果
			if i == 0 {
				result = node.Val
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

	// 遍历完所有层后，结果就为最后一层的第一个节点的值
	return result
}
