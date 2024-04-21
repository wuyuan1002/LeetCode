package main

// 872. 叶子相似的树

// 请考虑一棵二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个 叶值序列 。
// 如果有两棵二叉树的叶值序列是相同，那么我们就认为它们是 叶相似 的。
// 如果给定的两个根结点分别为 root1 和 root2 的树是叶相似的，则返回 true；否则返回 false 。

// leafSimilar .
// 深度优先遍历得到两棵树的叶子节点，然后比较两颗树的叶值序列是否一致
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// 构造两个列表用来存两棵数的叶值序列
	leaf1, leaf2 := make([]int, 0), make([]int, 0)

	// 获取两棵树的叶值序列
	dfsGetTreeLeaf(root1, &leaf1)
	dfsGetTreeLeaf(root2, &leaf2)

	// 比较两棵数的叶值序列是否完全一致
	if len(leaf1) != len(leaf2) {
		return false
	}
	for i := range leaf1 {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}
	return true
}

// dfsGetTreeLeaf 给定节点获取其所有叶子节点添加到结果中
func dfsGetTreeLeaf(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	// 若当前节点是叶子节点则将其值加入到结果中
	if node.Left == nil && node.Right == nil {
		*result = append(*result, node.Val)
	}

	// 遍历当前节点的左右子树
	dfsGetTreeLeaf(node.Left, result)
	dfsGetTreeLeaf(node.Right, result)
}
