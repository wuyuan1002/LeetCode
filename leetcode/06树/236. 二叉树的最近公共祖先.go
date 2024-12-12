package main

// 236. 二叉树的最近公共祖先

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，
// 最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

// lowestCommonAncestor236 .
// leetcode 235. 二叉搜索树的最近公共祖先
//
// 后序遍历二叉树，自底向上进行查找
// 分别在左右子树中查找p和q，第一个分别在左右子树中找到p和q的节点即为p和q的最近公共祖先
func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	// 递归到叶节点或当前节点为p或q，则返回当前节点
	if root == nil || root == p || root == q {
		return root
	}

	// 后序遍历二叉树，在左右子树中分别寻找p和q -- 返回值表示在树中找到的p或q或其最近公共祖先
	left := lowestCommonAncestor236(root.Left, p, q)
	right := lowestCommonAncestor236(root.Right, p, q)

	if left == nil {
		// 左子树没有找到p或q或其公共祖先节点，说明都在右子树中 -- 直接返回右子树的查询结果
		return right
	} else if right == nil {
		// 右子树没有找到p或q或其公共祖先节点，说明都在左子树中 -- 直接返回左子树的查询结果
		return left
	} else {
		// 左右子树结果都不为空，说明分别在左右子树中找到了p和q -- 当前节点即为最近的公共祖先节点
		return root
	}
}
