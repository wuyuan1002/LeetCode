package main

import (
	"fmt"
)

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。

func main() {
	fmt.Println(isValid(")"))
}

// 定义合法的括号匹配关系
var brackets = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

// 定义一个栈，遇到左括号则入栈，遇到右括号则和栈顶元素比较并弹出栈顶元素，若与栈顶元素不同，则说明不合法
func isValid(s string) bool {
	if s == "" {
		return true
	}

	stack := make([]byte, 0)

	sbyte := []byte(s)
	for _, c := range sbyte {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if (c == ')' || c == ']' || c == '}') && len(stack) > 0 {
			if c != brackets[stack[len(stack)-1]] {
				// 于栈中最后一个左括号对应的合法右括号比较，若不同则返回false
				return false
			}
			// 若当前右括号合法，则将栈顶左括号出栈，继续检查下一个右括号
			stack = stack[:len(stack)-1]
		} else {
			// 遇到了其他不合法的字符或栈中没有左括号却遇到了右括号
			return false
		}
	}

	return len(stack) == 0
}
