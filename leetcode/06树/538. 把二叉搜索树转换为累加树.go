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

// convertBST .
// 二叉搜索树左根右中序遍历是一个递增数组，那么倒过来进行右根左中序遍历，
// 遍历过程中进行累加，则每个节点的值就是递增数组当前节点值及其后的节点值累加
func convertBST(root *TreeNode) *TreeNode {
	val := 0
	dfsConvertBST(root, &val)
	return root
}

// dfsConvertBST 右根左中序遍历，同时进行累加
// preVal: 遍历的前一个节点的值
func dfsConvertBST(root *TreeNode, preVal *int) {
	if root == nil {
		return
	}

	// 遍历右子树
	dfsConvertBST(root.Right, preVal)

	// 将当前节点的值进行累加
	root.Val += *preVal
	// 标记当前节点为已遍历
	*preVal = root.Val

	// 遍历左子树
	dfsConvertBST(root.Left, preVal)
}
