package main

// 662. 二叉树最大宽度

// 给你一棵二叉树的根节点 root ，返回树的 最大宽度 。
//
// 树的 最大宽度 是所有层中最大的 宽度 。
//
// 每一层的 宽度 被定义为该层最左和最右的非空节点（即，两个端点）之间的长度。将这个二叉树视作与满二叉树结构相同，两端点间会出现一些延伸到这一层的 null 节点，这些 null 节点也计入长度。
//
// 题目数据保证答案将会在  32 位 带符号整数范围内。

// widthOfBinaryTree .
// 层级遍历二叉树（广度优先遍历），对每个节点进行编号，编号将树中的null节点也分配（即按照满二叉树的节点位置分配编号），
// 那么 最终某一层的宽度 = 最右侧节点编号 - 最左侧节点编号 + 1
//
// 满二叉树节点编号分配规则：若根节点编号为 u，则其左子节点编号为 u * 2，其右节点编号为 u * 2 + 1
func widthOfBinaryTree(root *TreeNode) int {

	result := 0                // 总结果
	queue := []pair{{root, 1}} // 使用一个单端队列存储遍历过程中的左右节点 -- 初始化先将root入队

	for queue != nil { // 若当前层有节点，则遍历该层

		// 获取当前层节点，并将队列清空用于存放下一层节点
		level := queue
		queue = nil

		// 计算当前层的宽度，并更新最大宽度 -- 某一层的宽度 = 最右侧节点编号 - 最左侧节点编号 + 1
		result = max(result, level[len(level)-1].index-level[0].index+1)

		// 遍历当前层节点，将节点入队列
		for _, p := range level {
			// 将当前节点的左右节点及其编号入队列 -- 为下一层的节点
			if p.node.Left != nil {
				queue = append(queue, pair{p.node.Left, p.index * 2})
			}
			if p.node.Right != nil {
				queue = append(queue, pair{p.node.Right, p.index*2 + 1})
			}
		}
	}

	return result
}

// pair 用来存节点及其编号
type pair struct {
	node  *TreeNode
	index int
}
