package main

// 617. 合并二叉树

// 给你两棵二叉树： root1 和 root2 。
//
// 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。
// 你需要将这两棵树合并成一棵新二叉树。
// 合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；
// 否则，不为 null 的节点将直接作为新二叉树的节点。
// 返回合并后的二叉树。
//
// 注意: 合并过程必须从两个树的根节点开始。

// mergeTrees .
// 前序遍历二叉树，先合并当前节点，之后递归合并子节点
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var merge *TreeNode // 合并后的节点

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
