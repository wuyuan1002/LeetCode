package main

//  重建二叉树

// 输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历 -- 第一个始终是根节点，后面部分的前半部分是它左子树前序遍历的结果，后半部分使它右子树前序遍历的结果
// 中序遍历 -- 根节点始终在中间，前半部分是它左子树中序遍历的结果，后半部分使它右子树中序遍历的结果

func main() {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}
	buildTree2(pre, in)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 { // 已经是叶子节点了
		return nil
	}

	leftPre := make([]int, 0)  // 左子树前序遍历数组
	rightPre := make([]int, 0) // 右子树前序遍历数组
	leftIn := make([]int, 0)   // 左子树中序遍历数组
	rightIn := make([]int, 0)  // 右子树中序遍历数组

	rootVal := preorder[0]     // 从前序遍历数组中获取根节点的值
	findedRootVal := false     // 是否找到根节点位置
	for val := range inorder { // 遍历中序遍历数组，构建左右子树的中序遍历数组
		if inorder[val] == rootVal {
			findedRootVal = true
			continue
		}
		if !findedRootVal {
			leftIn = append(leftIn, inorder[val])
		} else {
			rightIn = append(rightIn, inorder[val])
		}
	}

	leftPre = preorder[1 : len(leftIn)+1]
	rightPre = preorder[len(leftIn)+1:]

	root := &TreeNode{ // 创建根节点
		Val:   preorder[0],
		Left:  buildTree(leftPre, leftIn),
		Right: buildTree(rightPre, rightIn),
	}

	return root
}

// 第二次做
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 { // 已经是叶子节点了
		return nil
	}

	root := preorder[0] // 根节点
	index := 0          // 根节点在中序遍历数组中的下标
	for ; index < len(inorder); index++ {
		if inorder[index] == root {
			break
		}
	}

	leftIn := inorder[:index]              // 左子树的中序遍历数组
	rightIn := inorder[index+1:]           // 右子树的中序遍历数组
	leftPre := preorder[1 : len(leftIn)+1] // 左子树的前序遍历数组
	rightPre := preorder[len(leftIn)+1:]   // 右子树的前序遍历数组

	return &TreeNode{
		Val:   root,
		Left:  buildTree2(leftPre, leftIn),
		Right: buildTree2(rightPre, rightIn),
	}
}
