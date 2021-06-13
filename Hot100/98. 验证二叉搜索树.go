package main

import (
	"fmt"
	"math"
)

// 98. 验证二叉搜索树

// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征：
//
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。

func main() {
	aa := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(isValidBST(aa))
}

var pre = math.MinInt64

// 1. 二叉搜索树中序遍历后是排序的，因此，在中序遍历时校验当前值是否比前一个值大即可
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isValidBST(root.Left) {
		return false
	}

	if root.Val <= pre {
		return false
	}
	pre = root.Val

	if !isValidBST(root.Right) {
		return false
	}

	return true
}

// 2. 后序遍历，通过最大最小值限定，校验节点是否符合要求
func isValidBST1(root *TreeNode) bool {
	return isValidTree(root, math.MinInt64, math.MaxInt64)
}

// min、max: 当前节点的最小最大限定值
func isValidTree(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	// 左树所有节点的值都不能大于当前节点，右树所有节点都不能小于当前节点
	return isValidTree(root.Left, min, root.Val) && isValidTree(root.Right, root.Val, max)
}
