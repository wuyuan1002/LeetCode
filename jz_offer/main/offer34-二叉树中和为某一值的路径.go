package main

// 剑指 Offer 34. 二叉树中和为某一值的路径

// 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。
// 从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

// func main() {

// }

func pathSum(root *TreeNode, sum int) [][]int {
	allResult := make([][]int, 0)
	if root == nil {
		return allResult
	}

	result := make([]int, 0)
	currentSum := 0 // 现在路径上的数字总和
	getPath(root, sum, currentSum, result, &allResult)
	return allResult
}

func getPath(root *TreeNode, sum int, currentSum int, result []int, allResult *[][]int) {
	if root == nil {
		return
	}

	result = append(result, root.Val)
	currentSum += root.Val
	if currentSum == sum && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(result))
		copy(tmp, result)
		*allResult = append(*allResult, tmp)
	}

	getPath(root.Left, sum, currentSum, result, allResult)
	getPath(root.Right, sum, currentSum, result, allResult)

	// result = result[:len(result)-1] // result每次传递都是从父节点拷贝出来的值传递，如果result改成指针传递，则需要在返回父节点前把自己删掉
}

// 第二次做
func pathSum1(root *TreeNode, sum int) [][]int {
	allResult := make([][]int, 0)
	result := make([]int, 0)
	currentSum := 0

	getPath1(root, sum, currentSum, &result, &allResult)

	return allResult
}

func getPath1(root *TreeNode, sum int, currentSum int, result *[]int, allResult *[][]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	currentSum += root.Val

	if currentSum == sum && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(*result))
		copy(tmp, *result)
		*allResult = append(*allResult, tmp)
	}

	getPath1(root.Left, sum, currentSum, result, allResult)
	getPath1(root.Right, sum, currentSum, result, allResult)

	// 返回父节点前从路径中删掉自己
	*result = (*result)[:len(*result)-1]
}
