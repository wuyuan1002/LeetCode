package main

// 113. 路径总和 II

// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，
// 找出所有从根节点到叶子节点 路径总和等于给定目标和的路径。
//
// 叶子节点 是指没有子节点的节点。

func main() {

}

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	res := make([]int, 0)
	dfs3(root, 0, targetSum, &res, &result)

	return result
}

func dfs3(node *TreeNode, currentNum, targetSum int, res *[]int, result *[][]int) {
	if node == nil {
		return
	}

	currentNum += node.Val
	*res = append(*res, node.Val)

	if currentNum == targetSum && node.Left == nil && node.Right == nil {
		tmp := make([]int, len(*res))
		copy(tmp, *res)
		*result = append(*result, tmp)
	} else {
		dfs3(node.Left, currentNum, targetSum, res, result)
		dfs3(node.Right, currentNum, targetSum, res, result)
	}

	*res = (*res)[:len(*res)-1]
}
