package main

// 946. 验证栈序列

// 给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，只有当它们可能是在最初
// 空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false 。

// validateStackSequences .
// 使用一个栈进行模拟 压入/弹出 操作，若可以完成则说明没问题
func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}

	stack := make([]int, 0, len(pushed)) // 用于进行栈操作
	i := 0                               // 用于记录当前要弹出的数字下标

	for _, n := range pushed {
		// 将当前数字进行入栈
		stack = append(stack, n)

		// 循环判断栈顶元素与当前要弹出的数字是否相等，若相等则进行出栈操作
		for len(stack) > 0 && popped[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			i++
		}
	}

	return len(stack) == 0
}
