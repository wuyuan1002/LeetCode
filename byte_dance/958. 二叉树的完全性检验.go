package main

// 958. 二叉树的完全性检验

// 给定一个二叉树，确定它是否是一个完全二叉树。
//
// 百度百科中对完全二叉树的定义如下：
//
// 若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，
// 第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。（注：第 h 层可能包含 1~2h个节点。）

func main() {

}

// 层级遍历二叉树，若出现nil节点后又出现了节点，说明是非完全二叉树
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	end := false
	queue := []*TreeNode{root}
	node := root
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]

		if node == nil {
			end = true
		} else {
			if end {
				return false
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	return true
}
