package main

// 236. 二叉树的最近公共祖先

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，
// 最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

func main() {

}

// 后序遍历二叉树
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}

	// 访问到一个节点时:
	// 1.如果 p 和 q 都存在，则返回它们的公共祖先
	// 2.如果只存在一个，则返回存在的一个
	// 3.如果 p 和 q 都不存在，则返回nil

	// 求在左右子树中p和q的公共祖先节点
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		// 若左子树中没有p和q的公共祖先，说明p和q都不在左子树中，返回右子树的公共祖先
		return right
	} else if right == nil {
		// 若右子树中没有p和q的公共祖先，说明p和q都不在右子树中，返回左子树的公共祖先
		return left
	} else {
		// 若left和right都不为nil，说明p和q分别位于root的两侧，root为公共祖先节点
		return root
	}
}
