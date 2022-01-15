package main

// 剑指 Offer 31. 栈的压入、弹出序列

// 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。
// 假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，
// 序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

// func main() {
// 	pushed := []int{1, 2, 3, 4, 5}
// 	popped := []int{4, 5, 3, 2, 1}
// 	fmt.Printf("%v", validateStackSequences1(pushed, popped))
// }

func validateStackSequences(pushed []int, popped []int) bool {

	// isQueue:=false

	// 建一个辅助栈
	stack := make([]int, 0)
	// 将push顺序压入辅助栈中，如果栈顶元素==pop序列中下一个出现的值，则弹出
	i := 0
	for _, value := range popped {
		if len(stack) > 0 && value == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}
		if i >= len(pushed) {
			return false
		}
		for value != pushed[i] {
			stack = append(stack, pushed[i])
			i++
			if i >= len(pushed) {
				return false
			}
		}
		i++
	}
	return len(stack) == 0
}

// 第二次做 -- 建一个辅助栈，按照入栈序列依次入栈，直到遇到弹出序列的栈顶元素时，弹出元素
func validateStackSequences1(pushed []int, popped []int) bool {
	if pushed == nil && popped == nil {
		return true
	}
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}
	if pushed == nil || popped == nil {
		return false
	}

	stack := make([]int, 0)
	i := 0
	for _, val := range popped {
		if len(stack) > 0 && val == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}
		if i >= len(pushed) {
			return false
		}
		for val != pushed[i] {
			stack = append(stack, pushed[i])
			i++
			if i >= len(pushed) {
				return false
			}
		}
		i++
	}
	return len(stack) == 0
}
