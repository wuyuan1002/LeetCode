package main

// 617. 合并二叉树

// func main() {

// }

// 先序遍历二叉树，先合并当前节点，之后递归合并子节点
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	merge := (*TreeNode)(nil)
	if root1 == nil {
		merge = root2
	} else if root2 == nil {
		merge = root1
	} else {
		merge = &TreeNode{
			Val:   root1.Val + root2.Val,
			Left:  mergeTrees(root1.Left, root2.Left),
			Right: mergeTrees(root1.Right, root2.Right),
		}
	}

	return merge
}
