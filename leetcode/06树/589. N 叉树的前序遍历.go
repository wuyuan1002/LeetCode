package main

// 589. N 叉树的前序遍历

// 给定一个 n 叉树的根节点  root ，返回 其节点值的 前序遍历 。
//
// n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

// preorder .
func preorder(root *NodeNTree) []int {
	result := make([]int, 0)
	dfsPreOrder589(root, &result)
	return result
}

// dfsPreOrder589 深度优先前序遍历N叉树
func dfsPreOrder589(node *NodeNTree, result *[]int) {
	if node == nil {
		return
	}

	// 先将当前节点的值加入到结果集
	*result = append(*result, node.Val)

	// 然后遍历所有子节点
	for _, child := range node.Children {
		dfsPreOrder589(child, result)
	}
}
