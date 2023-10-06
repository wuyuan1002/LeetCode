package main

// 101. 对称二叉树

// 给你一个二叉树的根节点 root ， 检查它是否轴对称。

// isSymmetric .
// 同 Offer 28
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymme(root.Left, root.Right)
}

// isSymme 判断给定的两个节点是否镜像对称
func isSymme(node1, node2 *TreeNode) bool {

	// 判断node1和node2是否相等
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}

	// 判断node1与node2的子节点是否对称相等
	return isSymme(node1.Left, node2.Right) && isSymme(node1.Right, node2.Left)
}
