package main

// 700. 二叉搜索树中的搜索

// 给定二叉搜索树（BST）的根节点 root 和一个整数值 val。
//
// 你需要在 BST 中找到节点值等于 val 的节点。 返回以该节点为根的子树。
// 如果节点不存在，则返回 null 。

// searchBST 二叉搜索树中的搜索
// 前序遍历二叉树，若找到目标值则返回当前节点，否则递归遍历左右子树
func searchBST(root *TreeNode, val int) *TreeNode {
	// 当前节点为空或当前节点的值为目标值 -- 返回目标节点
	if root == nil || root.Val == val {
		return root
	}

	// 由于是二叉搜索树，所以小于当前节点的值都在左子树上，大于当前节点的值都在右子树上
	if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}

// searchBST1 普通二叉树中搜索
// 前序遍历二叉树，若找到目标值则返回当前节点，否则递归遍历左右子树
func searchBST1(root *TreeNode, val int) *TreeNode {
	// 当前节点为空或当前节点的值为目标值 -- 返回目标节点
	if root == nil || root.Val == val {
		return root
	}

	// 先去左子树找，左子树没找到再去右子树找
	result := searchBST1(root.Left, val)
	if result == nil {
		result = searchBST1(root.Right, val)
	}
	return result
}
