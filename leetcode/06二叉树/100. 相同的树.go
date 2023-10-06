package main

// 100. 相同的树

// 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
//
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

// isSameTree .
// 同 leetcode 101. 对称二叉树
func isSameTree(p *TreeNode, q *TreeNode) bool {

	// 判断p、q是否相等
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	// 判断p、q的左右子节点是否分别相等
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
