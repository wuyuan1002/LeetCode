package main

// 剑指 Offer 35. 复杂链表的复制

// 请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，
// 每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

// func main() {
// 	n7 := &Node{
// 		Val:    7,
// 		Next:   nil,
// 		Random: nil,
// 	}
// 	n13 := &Node{
// 		Val:    13,
// 		Next:   nil,
// 		Random: nil,
// 	}
// 	n11 := &Node{
// 		Val:    11,
// 		Next:   nil,
// 		Random: nil,
// 	}
// 	n10 := &Node{
// 		Val:    10,
// 		Next:   nil,
// 		Random: nil,
// 	}
// 	n1 := &Node{
// 		Val:    1,
// 		Next:   nil,
// 		Random: nil,
// 	}

// 	n7.Next = n13
// 	n7.Random = nil
// 	n13.Next = n11
// 	n13.Random = n7
// 	n11.Next = n10
// 	n11.Random = n1
// 	n10.Next = n1
// 	n10.Random = n11
// 	n1.Next = nil
// 	n1.Random = n7

// 	list2 := copyRandomList(n7)
// 	for node := list2; node != nil; node = node.Next {
// 		fmt.Println(node.Val)
// 	}

// }

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 自己写的，怎么调试都不好用 -- 不知道为什么，感觉没毛病啊
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	var next *Node
	for node := head; node != nil; node = next {
		next = node.Next
		node.Next = &Node{
			Val:    node.Val,
			Next:   node.Next,
			Random: node.Random,
		}
	}

	clonedHead := head.Next // 复制后的链表头
	for node := clonedHead; node != nil; node = node.Next {
		if node.Random != nil {
			node.Random = node.Random.Next
		}
	}

	for node := head; node != nil; node = next {
		next = node.Next
		if next != nil {
			node.Next = next.Next
		}
	}

	return clonedHead
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	p := head
	for p != nil {
		newNode := &Node{
			Val:    p.Val,
			Next:   nil,
			Random: nil,
		}
		newNode.Next = p.Next
		p.Next = newNode
		p = p.Next.Next
	}
	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}
	newHead := head.Next
	oldP := head
	p = newHead
	for p.Next != nil {
		oldP.Next = oldP.Next.Next
		p.Next = p.Next.Next
		oldP = oldP.Next
		p = p.Next
	}
	oldP.Next = nil
	return newHead
}
