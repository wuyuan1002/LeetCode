package main

// 337. 打家劫舍 III

// func main() {

// }

func rob2(root *TreeNode) int {
	rob, nrob := dfs19(root)
	return max(rob, nrob)
}

func dfs19(root *TreeNode) (rob int, nrob int) {
	if root == nil {
		return 0, 0
	}

	// 左右节点选择偷和不偷的钱数
	lrob, lnrob := dfs19(root.Left)
	rrob, rnrob := dfs19(root.Right)

	// 当前节点选择偷和不偷的钱数
	nrob = max(lrob, lnrob) + max(rrob, rnrob)
	rob = lnrob + rnrob + root.Val

	return rob, nrob
}
