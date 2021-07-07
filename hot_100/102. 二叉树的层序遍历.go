package main

// 102. 二叉树的层序遍历

// 给你一个二叉树，请你返回其按层序遍历得到的节点值。（即逐层地，从左到右访问所有节点）。

func main() {

}

// 广度优先遍历树
// 同offer 32
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0) // 总结果
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	num := 1     // 当前层个数
	nextNum := 0 // 下一层个数

	var node *TreeNode
	res := make([]int, 0) // 某一层的结果
	for i := 0; i < len(queue); i++ {
		node = queue[i]
		num--
		res = append(res, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
			nextNum++
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nextNum++
		}

		if num == 0 {
			result = append(result, res)
			res = make([]int, 0)
			num = nextNum
			nextNum = 0
		}
	}

	return result
}
