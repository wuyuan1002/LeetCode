package main

// 429. N 叉树的层序遍历

// 给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
//
// 树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

// levelOrder429 .
// 使用一个单端队列存储遍历过程中的所有子节点，不断遍历当前层节点将其下一层加入到队列中
func levelOrder429(root *NodeNTree) [][]int {
	result := make([][]int, 0)  // 结果集
	queue := []*NodeNTree{root} // 使用一个单端队列存储遍历过程中的所有子节点

	// 只要队列中有节点就说明没遍历完 -- 若当前层有节点，则遍历该层
	for len(queue) != 0 {
		// 当前层的结果，因为在遍历的过程中遍历过的节点会从queue中删除，所以遍历完一层后queue内的节点就是下一层的所有节点
		levelResult := make([]int, 0, len(queue))

		// 开始遍历当前层，count：当前层节点个数
		for count := len(queue); count > 0; count-- {
			// 获取当前层的下一个节点并出队列
			node := queue[0]
			queue = queue[1:]

			// 将该节点入结果集
			levelResult = append(levelResult, node.Val)

			// 将当前节点的子节点入队列 -- 为下一层的节点
			queue = append(queue, node.Children...)
		}

		// 遍历完当前层，将该层的结果入结果集，然后开始遍历下一层
		result = append(result, levelResult)
	}

	return result
}

// NodeNTree .
type NodeNTree struct {
	Val      int
	Children []*NodeNTree
}
