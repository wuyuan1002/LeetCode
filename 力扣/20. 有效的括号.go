package main

// 20. 有效的括号

// func main() {

// }

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
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, byte(c))
		} else if (c == ')' || c == ']' || c == '}') && len(stack) > 0 {
			if byte(c) != brackets[stack[len(stack)-1]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
