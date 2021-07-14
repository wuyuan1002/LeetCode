package main

// 226. 翻转二叉树

// 翻转一棵二叉树。

func main() {

}

// 同offer 27
// 前中后序遍历二叉树，把左右子树交换
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)

	root.Left, root.Right = root.Right, root.Left

	return root
}
