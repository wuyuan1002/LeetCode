package main

// 剑指 Offer 26. 树的子结构

// 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
// B是A的子结构， 即 A中有出现和B相同的结构和节点值。

func main() {

}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	have := false
	if A.Val == B.Val {
		have = doesTree1HasTree2(A, B)
	}
	if !have {
		have = isSubStructure(A.Left, B)
	}
	if !have {
		have = isSubStructure(A.Right, B)
	}
	return have
}

func doesTree1HasTree2(tree1 *TreeNode, tree2 *TreeNode) bool {
	if tree2 == nil {
		return true
	}
	if tree1 == nil {
		return false
	}
	if tree1.Val != tree2.Val {
		return false
	}
	return doesTree1HasTree2(tree1.Left, tree2.Left) && doesTree1HasTree2(tree1.Right, tree2.Right)
}

// 第二次做 -- 前序遍历大的树，找到值相同的根节点，再判断子树开始是否包含给定的树
func isSubStructure1(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	have := false
	if A.Val == B.Val {
		have = doestree1hastree21(A, B)
	}
	if !have {
		have = isSubStructure1(A.Left, B)
	}
	if !have {
		have = isSubStructure1(A.Right, B)
	}
	return have
}

func doestree1hastree21(tree1 *TreeNode, tree2 *TreeNode) bool {
	if tree2 == nil {
		return true
	}
	if tree1 == nil {
		return false
	}
	if tree1.Val != tree2.Val {
		return false
	}
	return doestree1hastree21(tree1.Left, tree2.Left) && doestree1hastree21(tree1.Right, tree2.Right)
}
