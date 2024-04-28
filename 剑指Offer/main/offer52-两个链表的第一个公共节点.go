package main

// 剑指 Offer 52. 两个链表的第一个公共节点

// ListNode 输入两个链表，找出它们的第一个公共节点。
type ListNode struct {
	Val  int
	Next *ListNode
}

// func main() {
// 	headA := &ListNode{
// 		Val: 4,
// 		Next: &ListNode{
// 			Val: 1,
// 			Next: &ListNode{
// 				Val: 8,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}

// 	headB := &ListNode{
// 		Val: 5,
// 		Next: &ListNode{
// 			Val: 0,
// 			Next: &ListNode{
// 				Val: 1,
// 				Next: &ListNode{
// 					Val: 8,
// 					Next: &ListNode{
// 						Val: 4,
// 						Next: &ListNode{
// 							Val:  5,
// 							Next: nil,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	node1 := getIntersectionNode1(headA, headB)
// 	fmt.Println(node1)
// }

// 方法1: 建立两个栈，将两个链表分别从头到尾遍历放入栈中，
// 之后从两个栈中同时弹出栈顶元素，直到两个栈顶元素不同时，则最后一个相同元素就是两个链表的第一个公共节点
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	stack1, stack2 := make([]*ListNode, 0), make([]*ListNode, 0)

	for headA != nil {
		stack1 = append(stack1, headA)
		headA = headA.Next
	}
	for headB != nil {
		stack2 = append(stack2, headB)
		headB = headB.Next
	}

	alen, blen := len(stack1)-1, len(stack2)-1
	var res *ListNode

	for alen >= 0 && blen >= 0 && stack1[alen] == stack2[blen] {
		res = stack1[alen]
		alen--
		blen--
	}

	return res
}

// 方法2: 分别遍历两个链表, 用map对第一个链表的节点做标记. 在对第二个链表遍历的时候,
// 判断每个节点是否已经被标记过. 如果有, 那么这个就是公共节点. 而链表是顺序表,
// 所以第一次碰到的公共节点一定是第一个公共节点, break跳出.
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool)

	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if m[headB] == true {
			break
		}
		headB = headB.Next
	}

	return headB
}

// 方法3: 双指针法
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	var temp1, temp2 *ListNode

	temp1 = headA
	temp2 = headB

	for temp1 != temp2 {
		if temp1 == nil {
			temp1 = headB
		} else {
			temp1 = temp1.Next
		}

		if temp2 == nil {
			temp2 = headA
		} else {
			temp2 = temp2.Next
		}
	}

	return temp1
}

// 第二次做
// 方法3: 双指针法
func getIntersectionNode4(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	nodea, nodeb := headA, headB
	for nodea != nodeb {
		if nodea != nil {
			nodea = nodea.Next
		} else {
			nodea = headB
		}

		if nodeb != nil {
			nodeb = nodeb.Next
		} else {
			nodeb = headA
		}
	}

	return nodea
}
