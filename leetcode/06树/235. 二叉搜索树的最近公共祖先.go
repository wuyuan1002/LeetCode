package main

// 235. 二叉搜索树的最近公共祖先

// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
// 满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

// lowestCommonAncestor235 .
// leetcode 236. 二叉树的最近公共祖先
//
// 1. 后序遍历自底向上查找 -- 同236
// 2. 前序遍历二叉搜索树自顶向下查找 -- 发现的第一个处于[p, q]范围的节点便是p和q的最近公共祖先
func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > p.Val && root.Val > q.Val {
		// 当前节点的值比p和q的值都大，说明p和q都在当前节点的左子树中 -- 去左子树继续查找
		return lowestCommonAncestor235(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		// 当前节点的值比p和q的值都小，说明p和q都在当前节点的右子树中 -- 去右子树继续查找
		return lowestCommonAncestor235(root.Right, p, q)
	} else {
		// 当前节点的值恰好位于 [p, q] 范围内 -- 当前节点为最近公共祖先
		return root
	}
}
