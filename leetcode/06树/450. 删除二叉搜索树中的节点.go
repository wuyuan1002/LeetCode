package main

// 450. 删除二叉搜索树中的节点

// 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，
// 并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
//
// 一般来说，删除节点可分为两个步骤：
//
// 首先找到需要删除的节点；
// 如果找到了，删除它。

// deleteNode .
// 1. 如果目标值大于当前节点值，则去右子树中删除；
// 2. 如果目标值小于当前节点值，则去左子树中删除；
// 3. 如果目标值为当前节点值，分为以下三种情况：
// 	3.1. 当前节点无左子树：其右子顶替其位置 -- 由此删除了该节点
// 	3.2. 当前节点无右子树：其左子顶替其位置 -- 由此删除了该节点
// 	3.3. 当前节点左右子树都有：其左子树转移到其右子树的最左节点的左子树上，然后右子树顶替其位置 -- 由此删除了该节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > key {
		// 当前节点值大于目标值 -- 去左子树删除
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		// 当前节点值小于目标值 -- 去右子树删除
		root.Right = deleteNode(root.Right, key)
	} else {
		// 当前节点为要被删除的节点

		if root.Left == nil {
			// 3.1. 当前节点没有左子树 -- 使用当前节点的右子树替换掉当前节点
			root = root.Right
		} else if root.Right == nil {
			// 3.2. 当前节点没有右子树 -- 使用当前节点的左子树替换掉当前节点
			root = root.Left
		} else {
			// 3.3. 当前节点左右子树都有

			// 寻找当前节点右子树的最左侧节点
			node := root.Right
			for node.Left != nil {
				node = node.Left
			}
			// 将当前节点的左子树转移到当前节点右子树的最左侧节点
			node.Left = root.Left
			// 将当前节点删除 -- 使用当前节点的右子树替换掉当前节点
			root = root.Right
		}

	}

	return root
}
