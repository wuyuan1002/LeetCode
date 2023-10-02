package main

// 572. 另一棵树的子树

// 给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot
// 具有相同结构和节点值的子树。如果存在，返回 true ；否则，返回 false 。
//
// 二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。
// tree 也可以看做它自身的一棵子树。

// func main() {

// }

// 类似 offer 26
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}

	is := false
	if root.Val == subRoot.Val {
		is = isSub(root, subRoot)
	}
	if !is {
		is = isSubtree(root.Left, subRoot)
	}
	if !is {
		is = isSubtree(root.Right, subRoot)
	}
	return is
}

func isSub(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil && root == nil {
		return true
	}
	if subRoot == nil || root == nil {
		return false
	}
	if root.Val != subRoot.Val {
		return false
	}
	return isSub(root.Left, subRoot.Left) && isSub(root.Right, subRoot.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
