package main

// 513. 找树左下角的值

// 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//
// 假设二叉树中至少有一个节点。

func main() {

}

// 层序遍历，找最后一层，左面第一个节点
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	result := -1
	isFirst := true // 是否为一层的第一个
	queue := []*TreeNode{root}
	num, nextNum := 1, 0 // 当前层剩余个数、下一层个数
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		num--
		if isFirst {
			result = node.Val
			isFirst = false
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			nextNum++
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nextNum++
		}

		if num == 0 {
			num = nextNum
			nextNum = 0
			isFirst = true
		}
	}

	return result
}
