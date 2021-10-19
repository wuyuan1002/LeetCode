package main

// 230. 二叉搜索树中第K小的元素

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

func main() {

}

// 中序遍历
func kthSmallest(root *TreeNode, k int) int {
	if k <= 0 {
		panic("必须大于0")
	}

	res := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		if k == 0 {
			// 已在左子树中找到目标节点
			return
		}
		if k--; k == 0 {
			// 当前节点为目标节点
			res = node.Val
			return
		}
		dfs(node.Right)
	}

	dfs(root)
	return res
}
