package main

// 剑指 Offer 33. 二叉搜索树的后序遍历序列

// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。
// 如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

// func main() {
// 	postorder := []int{4, 8, 6, 12, 16, 14, 10}
// 	// postorder := []int{1, 2, 5, 10, 6, 9, 4, 3}
// 	fmt.Printf("%v", verifyPostorder1(postorder))
// }

func verifyPostorder(postorder []int) bool {
	if postorder == nil || len(postorder) == 0 {
		return true
	}

	root := postorder[len(postorder)-1]                             // 根节点的值
	index := findLeftRightIndex(postorder[:len(postorder)-1], root) // 左右子树分界点
	leftTree := postorder[:index]                                   // 左数
	rightTree := postorder[index : len(postorder)-1]                // 右树

	if !(allMin(leftTree, root) && allMax(rightTree, root)) {
		return false
	}
	return verifyPostorder(leftTree) && verifyPostorder(rightTree)
}

func allMax(tree []int, root int) bool {
	if tree == nil || len(tree) == 0 {
		return true
	}
	for _, v := range tree {
		if v < root {
			return false
		}
	}
	return true
}

func allMin(tree []int, root int) bool {
	if tree == nil || len(tree) == 0 {
		return true
	}
	for _, v := range tree {
		if v >= root {
			return false
		}
	}
	return true
}

func findLeftRightIndex(postorder []int, root int) int {
	for i, value := range postorder {
		if value >= root {
			return i
		}
	}
	return len(postorder)
}

// 第二次做
func verifyPostorder1(postorder []int) bool {
	if postorder == nil || len(postorder) == 0 {
		return true
	}
	root := postorder[len(postorder)-1]
	index := findLeftRightIndex1(postorder[:len(postorder)-1], root)
	left := postorder[:index]
	right := postorder[index : len(postorder)-1]
	if !(allMin1(left, root) && allMax1(right, root)) {
		return false
	}
	return verifyPostorder1(left) && verifyPostorder1(right)
}

func findLeftRightIndex1(postorder []int, root int) int {
	for i, val := range postorder {
		if val >= root {
			return i
		}
	}
	return len(postorder)
}

func allMax1(postorder []int, root int) bool {
	for _, val := range postorder {
		if val < root {
			return false
		}
	}
	return true
}

func allMin1(postorder []int, root int) bool {
	for _, val := range postorder {
		if val >= root {
			return false
		}
	}
	return true
}
