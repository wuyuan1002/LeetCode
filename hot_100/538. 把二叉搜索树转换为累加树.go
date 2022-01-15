package main

// 538. 把二叉搜索树转换为累加树

// 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），
// 使每个节点 node的新值等于原树中大于或等于node.val的值之和。
//
// 提醒一下，二叉搜索树满足下列约束条件：
//
// 节点的左子树仅包含键 小于 节点键的节点。
// 节点的右子树仅包含键 大于 节点键的节点。
// 左右子树也必须是二叉搜索树。

// func main() {

// }

// 右根左中序遍历二叉树，遍历过程中进行累加，
// 因为二叉搜索树右节点始终大于当前节点，所以把右子树都累加起来再加上当前节点的值，就是当前节点的最新值，
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	convert(root)

	return root
}

var preSum = 0 // 遍历当前节点时，前一个节点累加到的和

func convert(root *TreeNode) {
	if root == nil {
		return
	}

	convert(root.Right)

	root.Val += preSum
	preSum = root.Val

	convert(root.Left)
}
