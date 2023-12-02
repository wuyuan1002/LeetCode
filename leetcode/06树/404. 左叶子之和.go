package main

// 404. 左叶子之和

// 给定二叉树的根节点 root ，返回所有左叶子之和。

// sumOfLeftLeaves .
// 前中后序遍历二叉树，遍历到左叶子节点时更新结果值
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	getSum(root.Left, true, &sum)
	getSum(root.Right, false, &sum)
	return sum
}

// getSum 获取左叶子之和，isLeft: 给定的节点node是否为其父节点的左节点
func getSum(node *TreeNode, isLeft bool, sum *int) {
	if node == nil {
		return
	}

	// 若当前节点为左叶子节点，则更新结果值
	if isLeft && node.Left == nil && node.Right == nil {
		*sum += node.Val
		return
	}

	// 继续递归当前节点的左右节点
	getSum(node.Left, true, sum)
	getSum(node.Right, false, sum)
}
