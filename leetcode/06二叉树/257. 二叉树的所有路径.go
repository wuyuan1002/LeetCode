package main

import "strconv"

// 257. 二叉树的所有路径

// 给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//
// 叶子节点 是指没有子节点的节点。

// binaryTreePaths .
// 递归遍历二叉树，遍历到叶节点时记录一条路径结果
func binaryTreePaths(root *TreeNode) []string {
	res := make([]int, 0)       // 一次遍历的路径节点
	result := make([]string, 0) // 总结果集

	dfsBinaryTreePaths(root, &res, &result)
	return result
}

// dfsBinaryTreePaths 回溯递归遍历节点获取路径
func dfsBinaryTreePaths(node *TreeNode, res *[]int, result *[]string) {
	if node == nil {
		return
	}

	// 将当前节点的值加入到本次遍历的路径中
	*res = append(*res, node.Val)

	if node.Left == nil && node.Right == nil {
		// 若当前节点为叶节点，则记录一条路径到总结果集
		path := ""
		for i, val := range *res {
			path += strconv.Itoa(val)
			if i < len(*res)-1 {
				path += "->"
			}
		}
		*result = append(*result, path)
	} else {
		// 若当前结果不是叶节点，则继续递归当前节点的左右节点
		dfsBinaryTreePaths(node.Left, res, result)
		dfsBinaryTreePaths(node.Right, res, result)
	}

	// 当前节点遍历完成，将当前节点的值移出本次遍历的路径中
	*res = (*res)[:len(*res)-1]
}
