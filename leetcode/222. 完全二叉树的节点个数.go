package main

// 222. 完全二叉树的节点个数

// 给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

func main() {

}

// 1. 没有利用上完全二叉树的特点 -- 遍历每一个节点进行求解
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// 2. 利用完全二叉树的特点进行求解 -- 使用满二叉树的公式求解
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftH, rightH := 0, 0 // 左有子树深度
	leftNode := root.Left
	rightNode := root.Right
	for leftNode != nil {
		leftNode = leftNode.Left
		leftH++
	}
	for rightNode != nil {
		rightNode = rightNode.Right
		rightH++
	}

	// 若左右子树深度相等，说明时满二叉树，直接使用公式求解即可
	if leftH == rightH {
		return (2 << leftH) - 1
	}
	// 若左右子树深度不相等，分别求当前节点的左右子树
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
