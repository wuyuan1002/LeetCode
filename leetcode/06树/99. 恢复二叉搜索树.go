package main

// 99. 恢复二叉搜索树

// 给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。
// 请在不改变其结构的情况下，恢复这棵树 。

// recoverTree .
//
// 二叉搜索树中序遍历是有序的，遍历过程中使用pre记录遍历到的前一个位置不正确的节点，
// 再遍历到第二个不正确的节点时将其与第一个进行交换
func recoverTree(root *TreeNode) {
	// pre用来记录中序遍历过程中的前一个节点用来和下一个节点比较大小，errorNode用来记录位置不正确的两个节点
	info := &Info{}

	// 中序遍历查找两个位置错误的节点
	dfsRecoverTree(root, info)

	// 若找到两个位置错误的节点则进行值的交换
	if info.errNode1 != nil && info.errNode2 != nil {
		info.errNode1.Val, info.errNode2.Val = info.errNode2.Val, info.errNode1.Val
	}
}

// dfsRecoverTree 中序遍历二叉树，查找两个位置错误的节点并将其记录下来
func dfsRecoverTree(node *TreeNode, info *Info) {
	if node == nil {
		return
	}

	// 遍历左子树
	dfsRecoverTree(node.Left, info)

	// 处理当前节点 -- 若当前节点位置不正确，则进行errorNode的赋值
	if info.pre != nil && info.pre.Val > node.Val {
		if info.errNode1 != nil {
			info.errNode2 = node
			// 剪枝 -- 若已经找到了两个错误节点则可以直接return没必要继续遍历其右子树
			return
		}
		info.errNode1, info.errNode2 = info.pre, node
	}

	// 记录遍历的前一个节点并遍历右子树
	info.pre = node
	dfsRecoverTree(node.Right, info)
}

// Info 用于保存在递归过程中的节点信息
type Info struct {
	pre      *TreeNode // 遍历前一个节点
	errNode1 *TreeNode // 第一个位置错误的节点
	errNode2 *TreeNode // 最后一个位置错误的节点
}
