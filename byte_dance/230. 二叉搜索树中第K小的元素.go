package main

// 230. 二叉搜索树中第K小的元素

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

func main() {

}

// 类似 Offer 54 二叉搜索树的第k大节点
func kthSmallest(root *TreeNode, k int) int {
	if root == nil || k <= 0 {
		return -1
	}

	dfs10(root, &k)
	return target
}

var target int

// 左中右 中序遍历
func dfs10(node *TreeNode, k *int) {
	if node == nil {
		return
	}

	dfs10(node.Left, k)

	if *k == 0 {
		return
	}

	*k--
	if *k == 0 {
		target = node.Val
		return
	}

	dfs10(node.Right, k)
}
