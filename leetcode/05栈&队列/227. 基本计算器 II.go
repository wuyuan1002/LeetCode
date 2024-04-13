package main

// 227. 基本计算器 II

// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
// 整数除法仅保留整数部分。
//
// 你可以假设给定的表达式总是有效的。所有中间结果将在 [-231, 231 - 1] 的范围内。
// 注意：不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。

// calculate .
// 加号：将数字压入栈 -- 相当于 +(num)
// 减号：将数字的相反数压入栈 -- 相当于 +(-num)
// 乘除号：计算数字与栈顶元素，并将栈顶元素替换为计算结果 -- 相当于 +(num / stack[len(stack)-1])、+(num * stack[len(stack)-1])
// 计算结束后累加站内所有数字即为结果
func calculate(s string) int {
	stack := make([]int, 0)
	preSign, num := '+', 0 // 当前数字的前一个符号、当前数字的值

	for i, c := range s {
		// 若当前位是数字，则累计当前数字的每一位 -- 如 793+32
		isDigit := '0' <= c && c <= '9'
		if isDigit {
			num = num*10 + int(c-'0')
		}

		// 若当前碰到了运算符，或执行到了最后一位，则触发运算
		if !isDigit && c != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}

			// 更新下一个数字前面的符号并归零数字方便记录下一个数字
			preSign = c
			num = 0
		}
	}

	// 累加栈内元素的值
	result := 0
	for _, n := range stack {
		result += n
	}

	return result
}
