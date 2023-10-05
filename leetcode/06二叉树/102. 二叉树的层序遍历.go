package main

// 102. 二叉树的层序遍历

// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

// levelOrder .
// 同 Offer 32
// 广度优先遍历二叉树 -- 遍历时使用一个队列存放遍历到的节点的左右节点
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0) // 结果集
	queue := []*TreeNode{root} // 使用一个队列存储遍历过程中的左右节点，初始化先将root入队

	for len(queue) > 0 { // 若当前层有值，则遍历该行
		levelResult := make([]int, 0, len(queue))     // 当前层的结果，因为在遍历的过程中遍历过的元素会从queue中删除，所以遍历完一层后queue内的元素就是下一层的所有元素
		for count := len(queue); count > 0; count-- { // 开始遍历当前层，count：当前层节点个数
			// 获取该层的下一个节点并出队列
			node := queue[0]
			queue = queue[1:]

			// 将该节点入结果集
			levelResult = append(levelResult, node.Val)

			// 将当前节点的左右节点入队列 -- 为下一层的节点
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 遍历完当前层，将该层的结果入结果集，然后开始遍历下一层
		result = append(result, levelResult)
	}

	return result
}
