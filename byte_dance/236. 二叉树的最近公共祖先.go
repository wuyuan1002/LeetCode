package main

// 236. 二叉树的最近公共祖先

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，
// 最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

// func main() {

// }

// 后序遍历二叉树
// 如果找到一个节点，发现左子树出现结点p，右子树出现节点q，
// 或者左子树出现结点q，右子树出现节点p，那么该节点就是节点p和q的最近公共祖先。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}

	// 在左右子树中寻找p或q
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		// 两个都在右子树中
		return right
	} else if right == nil {
		// 两个都在左子树中
		return left
	} else {
		// root为公共祖先
		return root
	}
}
