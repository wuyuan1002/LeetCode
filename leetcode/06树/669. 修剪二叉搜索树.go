package main

// 669. 修剪二叉搜索树

// 给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。
// 通过修剪二叉搜索树，使得所有节点的值在[low, high]中。
// 修剪树 不应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。
// 可以证明，存在 唯一的答案 。
//
// 所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。

// trimBST .
// 前序遍历二叉树
// 1. 若当前节点值比要求的最小值都小，说明当前节点及其左子树全部应该被修剪掉，直接返回修建右子树的结果
// 2. 若当前节点值比要求的最大值都大，说明当前节点及其右子树全部应该被修剪掉，直接返回修建左子树的结果
// 3. 若当前节点的值位于目标区间内，则分别修剪当前节点的左右子树
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		// 当前节点值比要求的最小值都小，说明当前节点及其左子树全部应该被修剪掉，直接返回修建右子树的结果
		return trimBST(root.Right, low, high)
	} else if root.Val > high {
		// 当前节点值比要求的最大值都大，说明当前节点及其右子树全部应该被修剪掉，直接返回修建左子树的结果
		return trimBST(root.Left, low, high)
	} else {
		// 当前节点的值位于目标区间内，则分别修剪当前节点的左右子树
		root.Left = trimBST(root.Left, low, high)
		root.Right = trimBST(root.Right, low, high)
		return root
	}
}
