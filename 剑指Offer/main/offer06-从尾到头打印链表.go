package main

// 剑指 Offer 06. 从尾到头打印链表
// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

// func main() {

// }

// 递归法 -- 始终先遍历节点的后一个元素知道最后一个节点，把节点迅如数组中
func reversePrint1(head *ListNode) []int {
	arr := make([]int, 0)
	digui(head, &arr)
	return arr
}

func digui(head *ListNode, arr *[]int) {
	if head == nil {
		return
	}
	digui(head.Next, arr)
	*arr = append(*arr, head.Val)
}

// 2.先将节点挨个放入栈中，然后出栈放入数组中
func reversePrint2(head *ListNode) []int {
	var arr []int
	arrFan := make([]int, 0)
	for ; head != nil; head = head.Next {
		arr = append(arr, head.Val)
	}
	// 将栈中元素放入数组
	for i := range arr {
		arrFan = append(arrFan, arr[len(arr)-1-i])
	}
	return arrFan
}

// 3.先将链表元素挨个放入数组中，再将数组元素反转
func reversePrint3(head *ListNode) []int {
	arr := make([]int, 0)
	for ; head != nil; head = head.Next {
		arr = append(arr, head.Val)
	}
	// 反转数组
	for i, j := 0, len(arr)-1; i < j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}
