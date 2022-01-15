package main

// 99. 恢复二叉搜索树

// 给你二叉搜索树的根节点 root ，该树中的两个节点被错误地交换。请在不改变其结构的情况下，恢复这棵树。
//
// 进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用常数空间的解决方案吗？

// func main() {

// }

// 1.中序遍历二叉树放入数组中，之后遍历数组，交换两个位置不正确的数字位置
// 2.二叉搜索树中序遍历是有序的，遍历过程中使用pre记录遍历到的前一个节点
func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	// 需交换的节点:
	// 第一个节点s1，是第一个按照中序遍历时候前一个节点大于后一个节点，我们选取前一个节点
	// 第二个节点s2，是在第一个节点找到之后，后面出现前一个节点大于后一个节点，我们选择后一个节点
	var pre, s1, s2 *TreeNode

	// 定义中序遍历函数
	var dfs20 func(node *TreeNode)
	dfs20 = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs20(node.Left)

		if pre != nil && pre.Val > node.Val {
			if s1 != nil {
				s2 = node
				return
			}
			s1, s2 = pre, node
		}

		pre = node
		dfs20(node.Right)
	}

	// 遍历根节点
	dfs20(root)
	if s1 != nil && s2 != nil {
		s1.Val, s2.Val = s2.Val, s1.Val
	}
}
