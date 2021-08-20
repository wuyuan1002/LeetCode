package main

import (
	"container/list"
)

// 103. 二叉树的锯齿形层序遍历

// 给定一个二叉树，返回其节点值的锯齿形层序遍历。
// （即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

func main() {

}

// 广度优先遍历二叉树
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)

	queue := list.New()
	queue.PushBack(root)

	k := 1          // 当前行数
	currentNum := 1 // 当前行总数

	var node *TreeNode
	res := make([]int, 0)
	for currentNum > 0 {
		if k%2 == 1 {
			node = queue.Remove(queue.Front()).(*TreeNode)
			res = append(res, node.Val)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		} else {
			node = queue.Remove(queue.Back()).(*TreeNode)
			res = append(res, node.Val)

			if node.Right != nil {
				queue.PushFront(node.Right)
			}
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
		}
		currentNum--

		if currentNum == 0 {
			result = append(result, res)
			res = make([]int, 0)
			currentNum = queue.Len()
			k++
		}
	}

	return result
}

func zigzagLevelOrder1(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := list.List{}
	queue.PushBack(root)
	currentNum := 1
	n := 1
	res := make([]int, 0)
	for currentNum > 0 {
		if n%2 == 1 {
			node := queue.Remove(queue.Front()).(*TreeNode)
			currentNum--
			res = append(res, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		} else {
			node := queue.Remove(queue.Back()).(*TreeNode)
			currentNum--
			res = append(res, node.Val)
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
		}

		if currentNum == 0 {
			result = append(result, res)
			res = make([]int, 0)
			currentNum = queue.Len()
			n++
		}
	}

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
