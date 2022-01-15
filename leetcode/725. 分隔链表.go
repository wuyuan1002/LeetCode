package main

// 725. 分隔链表

// 给你一个头结点为 head 的单链表和一个整数 k ，请你设计一个算法将链表分隔为 k 个连续的部分。
// 每部分的长度应该尽可能的相等：任意两部分的长度差距不能超过 1 。这可能会导致有些部分为 null 。
// 这 k 个部分应该按照在链表中出现的顺序排列，并且排在前面的部分的长度应该大于或等于排在后面的长度。
// 返回一个由上述 k 部分组成的数组。

// func main() {

// }

// 先遍历一次数组得到链表长度，再计算每一段的长度，将结果放入结果集中
func splitListToParts(head *ListNode, k int) []*ListNode {
	num := 0 // 链表长度
	for node := head; node != nil; node = node.Next {
		num++
	}
	avg, remain := num/k, num%k

	parts := make([]*ListNode, k)
	for i, node := 0, head; i < k && node != nil; i++ {
		parts[i] = node
		size := avg // 一组的节点数量
		if i < remain {
			size++
		}
		for j := 1; j < size; j++ {
			node = node.Next
		}
		node, node.Next = node.Next, nil
	}
	return parts
}
