package main

// 剑指 Offer 28. 对称的二叉树

// 请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

// func main() {

// }

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSymmetricHelper(p.Left, q.Right) && isSymmetricHelper(p.Right, q.Left)
}

// 第二次做
func isSymmetric1(root *TreeNode) bool {
	return isSymmetricHelper1(root, root)
}

func isSymmetricHelper1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSymmetricHelper1(p.Left, q.Right) && isSymmetricHelper1(p.Right, q.Left)
}
