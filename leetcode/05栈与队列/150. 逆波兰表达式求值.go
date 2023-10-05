package main

import "strconv"

// 150. 逆波兰表达式求值

// 给你一个字符串数组 tokens ，表示一个根据 逆波兰表示法 表示的算术表达式。
// 请你计算该表达式。返回一个表示表达式值的整数。
//
// 注意：
//
// 有效的算符为 '+'、'-'、'*' 和 '/' 。
// 每个操作数（运算对象）都可以是一个整数或者另一个表达式。
// 两个整数之间的除法总是 向零截断 。
// 表达式中不含除零运算。
// 输入是一个根据逆波兰表示法表示的算术表达式。
// 答案及所有中间计算结果可以用 32 位 整数表示。

// evalRPN .
// 遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		if val, err := strconv.Atoi(token); err == nil {
			// 遇到数字则入栈
			stack = append(stack, val)
		} else {
			// 遇到操作符则取出栈顶两元素进行计算后将结果入栈
			num1, num2 := stack[len(stack)-2], stack[(len(stack))-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}

	// 最终的栈顶元素即为最终结果
	return stack[0]
}
