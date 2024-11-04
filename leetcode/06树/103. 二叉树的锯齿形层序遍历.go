package main

import "container/list"

// 103. 二叉树的锯齿形层序遍历

// 给定一个二叉树，返回其节点值的锯齿形层序遍历。
// （即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

// zigzagLevelOrder .
// leetcode 102. 二叉树的层序遍历
// Offer 32 - 从上到下打印二叉树
//
// 广度优先遍历二叉树
// 使用一个双端队列存放遍历到的节点的左右节点，不断遍历该队列直至所有节点都被遍历完
// 遍历到偶数行时，从队头读，子节点按照左右顺序从队尾入队
// 遍历到奇数行时，从队尾读，子节点按照右左顺序从队头入队
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0) // 总结果集
	queue := list.New()        // 使用一个双端队列存储遍历过程中的左右节点 -- 初始化先将root入队
	queue.PushBack(root)

	// 队列不为空说明还有层没有遍历，不断遍历每一层，同时使用levelNum记录当前层的行号
	for levelNum := 0; queue.Len() > 0; levelNum++ {

		// 当前层的结果
		levelResult := make([]int, 0, queue.Len())

		// 开始遍历当前层，此时队列内的所有节点都是当前层的节点
		// 因为在遍历的过程中遍历过的节点会从queue中删除，所以遍历完一层后队列内的节点就是下一层的所有节点
		for count := queue.Len(); count > 0; count-- {

			// 偶数行 -- 从队头读，子节点按照左右顺序从队尾入队
			// 奇数行 -- 从队尾读，子节点按照右左顺序从队头入队
			if levelNum%2 == 0 {
				// 从队头获取当前层的下一个节点并出队列
				node := queue.Remove(queue.Front()).(*TreeNode)

				// 将当前节点入结果集
				levelResult = append(levelResult, node.Val)

				// 将当前节点的子节点按照左右顺序从队尾入队 -- 为下一层的节点
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
				if node.Right != nil {
					queue.PushBack(node.Right)
				}
			} else {
				// 从队尾获取当前层的下一个节点并出队列
				node := queue.Remove(queue.Back()).(*TreeNode)

				// 将当前节点入结果集
				levelResult = append(levelResult, node.Val)

				// 当前节点的子节点按照右左顺序从队头入队 -- 为下一层的节点
				if node.Right != nil {
					queue.PushFront(node.Right)
				}
				if node.Left != nil {
					queue.PushFront(node.Left)
				}
			}
		}

		// 将当前行的结果入结果集
		result = append(result, levelResult)
	}

	return result
}
