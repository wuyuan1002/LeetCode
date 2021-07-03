package main

// 101. 对称二叉树

// 给定一个二叉树，检查它是否是镜像对称的。

func main() {

}

func isSymmetric(root *TreeNode) bool {
	return isSymme(root, root)
}

func isSymme(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}

	return isSymme(root1.Left, root2.Right) && isSymme(root1.Right, root2.Left)
}
