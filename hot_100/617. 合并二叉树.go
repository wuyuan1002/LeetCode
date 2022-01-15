package main

// 617. 合并二叉树

// 给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
//
// 你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，
// 那么将他们的值相加作为节点合并后的新值，否则不为NULL 的节点将直接作为新二叉树的节点。

// func main() {

// }

// 先序遍历二叉树，先合并当前节点，之后递归合并子节点
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	var merge *TreeNode // 合并后的树
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
