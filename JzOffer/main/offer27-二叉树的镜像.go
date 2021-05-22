package main

// 剑指 Offer 27. 二叉树的镜像

// 请完成一个函数，输入一个二叉树，该函数输出它的镜像。

func main() {

}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	root.Left = mirrorTree(root.Left)
	root.Right = mirrorTree(root.Right)
	return root
}

// 第二次做 -- 前中后序遍历树，把每一个节点的左右节点交换
func mirrorTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	mirrorTree1(root.Left)
	mirrorTree1(root.Right)

	root.Left, root.Right = root.Right, root.Left

	return root
}
