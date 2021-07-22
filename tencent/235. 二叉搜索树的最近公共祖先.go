package main

// 235. 二叉搜索树的最近公共祖先

// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，
// 最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

func main() {

}

// 二叉搜索树是排序的，从上到下第一个位于两个值中间的数字就是最近公共祖先
// 前序遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	min := min(p.Val, q.Val)
	max := max(p.Val, q.Val)

	if root.Val > max {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < min {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
