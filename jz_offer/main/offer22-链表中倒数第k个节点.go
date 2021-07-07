package main

// 剑指 Offer 22. 链表中倒数第k个节点
// 输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，
// 本题从1开始计数，即链表的尾节点是倒数第1个节点。
// 例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。
// 这个链表的倒数第 3 个节点是值为 4 的节点。

func main() {

}

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func getKthFromEnd(head *ListNode2, k int) *ListNode2 {
	if head == nil || k <= 0 {
		return nil
	}
	i, j := head, head
	for num := 0; num < k-1; num++ {
		if i.Next != nil {
			i = i.Next
		} else {
			return nil
		}
	}
	for i.Next != nil {
		i = i.Next
		j = j.Next
	}
	return j
}

// 第二次做 --
// 1.可以定义一个栈，先都入栈，再取值
// 2.一直递归到最后一个，到倒数第k个时输出结果。这样做的话，链表有几个节点，就需要递归几层，而且需要全局变量
// 3.定义两个指针，一个指针先走k-1步，之后两个指针同时向后走，当前面指针为nil时，后一个指针就正好是倒数第k个
func getKthFromEnd1(head *ListNode2, k int) *ListNode2 {
	return nil
}
