package main

// 105. 从前序与中序遍历序列构造二叉树

// 根据一棵树的前序遍历与中序遍历构造二叉树。
// 注意:
// 你可以假设树中没有重复的元素。

func main() {

}

// 同offer 07
func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 {
		return nil
	}
	if inorder == nil || len(inorder) == 0 {
		return nil
	}

	// 当前的根节点
	rootVal := preorder[0]

	// 当前根节点在中序遍历中的下标 -- 根节点左子树的节点个数
	rootIndex := 0
	for ; rootIndex < len(inorder); rootIndex++ {
		if inorder[rootIndex] == rootVal {
			break
		}
	}

	// 左子树前序遍历序列
	leftPre := preorder[1 : rootIndex+1]
	// 右子树前序遍历序列
	rightPre := preorder[rootIndex+1:]
	// 左子树中序遍历序列
	leftIn := inorder[:rootIndex]
	// 右子树中序遍历序列
	rightIn := inorder[rootIndex+1:]

	// 构建二叉树
	node := &TreeNode{
		Val:   rootVal,
		Left:  buildTree(leftPre, leftIn),
		Right: buildTree(rightPre, rightIn),
	}

	return node
}
