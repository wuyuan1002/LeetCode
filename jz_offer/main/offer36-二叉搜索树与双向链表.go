package main

/**
 * 剑指 Offer 36. 二叉搜索树与双向链表
 *
 * 输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。
 * 要求不能创建任何新的节点，只能调整树中节点指针的指向。
 *
 * 左指针表示双链表向前指，右指针表示双链表向后指
 */

func main() {

}

var (
	pre  *TreeNode = nil // 记录遍历的上一个节点
	head *TreeNode = nil // 记录链表头节点
)

// 类似 114
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	convert(root)

	pre.Right = head
	head.Left = pre

	return head
}

func convert(node *TreeNode) {
	if node == nil {
		return
	}
	convert(node.Left)

	// 二叉搜索树中序遍历后是有序的 -- 前一个节点就是已经访问过的排好序的最大值，当前根节点是第一个比它大的值
	// 每次到根节点，把根节点与记录的前一个节点连接，并把当前根节点记录为前一个结点
	if pre == nil {
		head = node
	} else {
		pre.Right = node
		node.Left = pre
	}
	pre = node
	convert(node.Right)
}

// 第二次做
func treeToDoublyList1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	convert1(root)
	pre.Right = head
	head.Left = pre
	return head
}

func convert1(node *TreeNode) {
	if node == nil {
		return
	}

	convert1(node.Left)
	if pre == nil {
		head = node
	} else {
		pre.Right = node
		node.Left = pre
	}
	pre = node
	convert1(node.Right)
}
