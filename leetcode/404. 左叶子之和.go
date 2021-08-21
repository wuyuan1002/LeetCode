package main

// 404. 左叶子之和

// 计算给定二叉树的所有左叶子之和。

func main() {

}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	getSum(root.Left, true, &sum)
	getSum(root.Right, false, &sum)
	return sum
}

func getSum(node *TreeNode, isLeft bool, sum *int) {
	if node == nil {
		return
	}

	if isLeft && node.Left == nil && node.Right == nil {
		*sum += node.Val
		return
	}

	getSum(node.Left, true, sum)
	getSum(node.Right, false, sum)
}
