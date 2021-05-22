package main

import (
	"container/list"
	"fmt"
)

// 剑指 Offer 32 - III. 从上到下打印二叉树 III

// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，
// 第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

func main() {
	root := &TreeNode10{
		Val: 1,
		Left: &TreeNode10{
			Val: 2,
			Left: &TreeNode10{
				Val: 4,
			},
		},
		Right: &TreeNode10{
			Val: 3,
			Right: &TreeNode10{
				Val: 5,
			},
		},
	}
	fmt.Printf("%v", levelOrder3(root))
}

type TreeNode10 struct {
	Val   int
	Left  *TreeNode10
	Right *TreeNode10
}

func levelOrder3(root *TreeNode10) [][]int {
	result := make([][]int, 0)
	if root != nil {
		queue := list.New()
		queue.PushBack(root)
		var levelResult []int

		k := 1 // 表示层数
		n := 1 // 当前行还有几个
		var node *TreeNode10
		for n > 0 {
			if k%2 == 1 {
				node = queue.Remove(queue.Front()).(*TreeNode10)
				levelResult = append(levelResult, node.Val)
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
				if node.Right != nil {
					queue.PushBack(node.Right)
				}
			} else if k%2 == 0 {
				node = queue.Remove(queue.Back()).(*TreeNode10)
				levelResult = append(levelResult, node.Val)
				if node.Right != nil {
					queue.PushFront(node.Right)
				}
				if node.Left != nil {
					queue.PushFront(node.Left)
				}
			}
			n--
			if n == 0 {
				n = queue.Len()
				result = append(result, levelResult)
				k++
				levelResult = make([]int, 0)
			}
		}
	}
	return result
}

// 第二次做 -- 使用两个栈来保存
