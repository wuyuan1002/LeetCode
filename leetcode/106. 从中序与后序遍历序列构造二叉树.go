package main

// 106. 从中序与后序遍历序列构造二叉树

// 根据一棵树的中序遍历与后序遍历构造二叉树。
//
// 注意:
// 你可以假设树中没有重复的元素。

func main() {

}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || len(inorder) == 0 || postorder == nil || len(postorder) == 0 {
		return nil
	}

	val := postorder[len(postorder)-1] // 当前节点的值
	index := 0                         // 当前节点在中序遍历中的下标
	for ; inorder[index] != val; index++ {
	}

	linorder := inorder[:index]
	rinorder := inorder[index+1:]
	lpost := postorder[:len(linorder)]
	rpost := postorder[len(lpost) : len(postorder)-1]

	node := &TreeNode{
		Val:   val,
		Left:  buildTree(linorder, lpost),
		Right: buildTree(rinorder, rpost),
	}
	return node
}
