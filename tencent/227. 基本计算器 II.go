package main

import (
	"fmt"
)

// 227. 基本计算器 II

// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
//
// 整数除法仅保留整数部分。

func main() {
	fmt.Println(calculate("37+2*3"))
}

// 加号：将数字压入栈；
// 减号：将数字的相反数压入栈；
// 乘除号：计算数字与栈顶元素，并将栈顶元素替换为计算结果。
func calculate(s string) int {
	if s == "" {
		return -1
	}

	stack := make([]int, 0)
	preSign := '+' // 当前数字的前一个符号
	num := 0
	for i, c := range s {
		isDigit := '0' <= c && c <= '9'
		if isDigit {
			num = num*10 + int(c-'0')
		}
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
			preSign = c
			num = 0
		}
	}

	// 类即栈内元素的值
	res := 0
	for _, n := range stack {
		res += n
	}

	return res
}
