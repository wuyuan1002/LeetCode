package main

// 给定一个二叉树，检查它是否是镜像对称的。
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

func main() {

}

// 同offer 28
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymme(root, root)
}

func isSymme(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	return isSymme(node1.Left, node2.Right) && isSymme(node1.Right, node2.Left)
}
