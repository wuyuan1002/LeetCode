package main

// 剑指 Offer 54. 二叉搜索树的第k大节点

// 给定一棵二叉搜索树，请找出其中第k大的节点。

// func main() {

// }

func kthLargest(root *TreeNode, k int) int {
	if root == nil || k <= 0 {
		return -1
	}

	find(root, &k)

	return target
}

var target int

func find(root *TreeNode, k *int) {
	if root == nil {
		return
	}

	find(root.Right, k)

	if *k == 0 { // 若已经找到，则直接返回即可，不用再继续遍历左子树
		return
	}

	*k--
	if *k == 0 {
		target = root.Val
		return
	}

	find(root.Left, k)
}

// 第二次做
func find1(root *TreeNode, k *int) {
	if root == nil {
		return
	}

	find1(root.Right, k)

	if *k == 0 {
		return
	}

	*k--
	if *k == 0 {
		target = root.Val
		return
	}

	find1(root.Left, k)
}
