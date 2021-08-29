package main

import (
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

}

func isValidBST(root *TreeNode) bool {
	return check(root, math.MinInt64, math.MaxInt64)
}

// min、max: 当前节点的最小最大限定值
func check(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	return check(node.Left, min, node.Val) && check(node.Right, node.Val, max)
}

// 中序遍历 -- 若为二叉搜索树则中序遍历为递增的
var preVal = math.MinInt32 // 前一个访问的节点值

func isValidBST1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isValidBST1(root.Left) {
		return false
	}

	if preVal > root.Val {
		return false
	}
	preVal = root.Val

	return isValidBST1(root.Right)
}
