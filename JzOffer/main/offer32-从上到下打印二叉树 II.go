package main

import (
	"container/list"
)

// 剑指 Offer 32 - II. 从上到下打印二叉树 II

// 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

func main() {

}

func levelOrder1(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root != nil {
		queue := []*TreeNode{root}
		var node *TreeNode
		nextLevel := 0                // 下一行的个数
		toBePrint := 1                // 当前行还有几个
		levelResult := make([]int, 0) // 存放一行的所有值
		for i := 0; i < len(queue); i++ {
			node = queue[i]
			levelResult = append(levelResult, node.Val)
			toBePrint--

			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevel++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevel++
			}

			if toBePrint == 0 { // 一行打印完成
				result = append(result, levelResult)
				levelResult = make([]int, 0) // 新建一个行结果切片
				toBePrint = nextLevel        // 开始打印下一行
				nextLevel = 0
			}
		}
	}
	return result
}

func levelOrder2(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root != nil {
		nextLevelQueue := []*TreeNode{root} // 存放一行的所有节点
		var levelResult []int               // 存放一行的所有值

		queue := nextLevelQueue
		var node *TreeNode
		for len(queue) > 0 { // 当前行有值，则下面去遍历该行
			nextLevelQueue = make([]*TreeNode, 0) // 存放下一行的元素
			levelResult = make([]int, 0)          // 存放当前行的值
			for i := 0; i < len(queue); i++ {     // 遍历当前行
				node = queue[i]
				levelResult = append(levelResult, node.Val)
				if node.Left != nil {
					nextLevelQueue = append(nextLevelQueue, node.Left)
				}
				if node.Right != nil {
					nextLevelQueue = append(nextLevelQueue, node.Right)
				}
			}

			result = append(result, levelResult) // 把当前行结果存到总结果中
			queue = nextLevelQueue               // 遍历下一行
		}
	}
	return result
}

// 第二次做 -- 在32的基础上，每打印一行多打印一个换行，添加一个计数器计算打印的个数
func levelOrder5(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0) // 所有行结果
	levelRes := make([]int, 0) // 一行的结果
	queue := list.New()
	queue.PushBack(root)

	num := 1     // 当前行还有几个
	nextNum := 0 // 下一行的个数

	var node *TreeNode
	for queue.Len() > 0 {
		node = queue.Remove(queue.Front()).(*TreeNode)
		num--
		levelRes = append(levelRes, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
			nextNum++
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
			nextNum++
		}

		if num == 0 { // 当前行已经打印完了
			result = append(result, levelRes)
			levelRes = make([]int, 0) // 新建一个行结果切片
			num = nextNum             // 开始打印下一行
			nextNum = 0
		}
	}
	return result
}
