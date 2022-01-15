package main

// 105. 从前序与中序遍历序列构造二叉树

// 根据一棵树的前序遍历与中序遍历构造二叉树。
//
// 注意:
// 你可以假设树中没有重复的元素。

// func main() {

// }

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 {
		return nil
	}
	if inorder == nil || len(inorder) == 0 {
		return nil
	}

	root := preorder[0] // 当前根节点的值
	rootIndex := 0      // 根节点在中序遍历中的下标
	for ; rootIndex < len(inorder); rootIndex++ {
		if inorder[rootIndex] == root {
			break
		}
	}

	leftPre := preorder[1 : rootIndex+1] // 左树的前序遍历序列
	rightPre := preorder[rootIndex+1:]   // 右树的前序遍历序列
	leftIn := inorder[:rootIndex]        // 左树的中序遍历序列
	rightIn := inorder[rootIndex+1:]     // 右树的中序遍历序列

	return &TreeNode{
		Val:   root,
		Left:  buildTree(leftPre, leftIn),
		Right: buildTree(rightPre, rightIn),
	}
}
