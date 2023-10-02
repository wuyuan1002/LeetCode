package main

import (
	"strconv"
	"strings"
)

// 257. 二叉树的所有路径

// 给定一个二叉树，返回所有从根节点到叶子节点的路径。
//
// 说明: 叶子节点是指没有子节点的节点。

// func main() {

// }

// 回溯法
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	result := make([]string, 0)

	dfs11(root, &res, &result)
	return result
}

func dfs11(root *TreeNode, res *[]int, result *[]string) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)

	if root.Left == nil && root.Right == nil {
		str := strings.Builder{}
		for i, n := range *res {
			str.WriteString(strconv.Itoa(n))
			if i != len(*res)-1 {
				str.WriteString("->")
			}
		}

		*result = append(*result, str.String())
	} else {
		dfs11(root.Left, res, result)
		dfs11(root.Right, res, result)
	}

	*res = (*res)[:len(*res)-1]
}
