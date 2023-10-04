package main

// 20. 有效的括号

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

// isValid .
// 定义一个栈，遇到左括号则入栈，遇到右括号则和栈顶元素比较并弹出栈顶元素，若与栈顶元素不同，则说明不合法
func isValid(s string) bool {
	// 定义合法的括号匹配关系
	brackets := map[byte]byte{'(': ')', '[': ']', '{': '}'}

	// 定义一个栈用来存左括号
	stack := make([]byte, 0)

	// 遍历字符串的每一个字符
	for _, c := range []byte(s) {
		if c == '(' || c == '[' || c == '{' {
			// 遇到左括号则入栈
			stack = append(stack, c)
		} else if len(stack) > 0 && (c == ')' || c == ']' || c == '}') {
			if c != brackets[stack[len(stack)-1]] {
				// 与栈中最后一个左括号对应的合法右括号比较，若不同则返回false
				return false
			}
			// 若当前右括号合法，则将栈顶左括号出栈，继续检查下一个字符
			stack = stack[:len(stack)-1]
		} else {
			// 遇到了其他不合法的字符或栈中没有左括号却遇到了右括号
			return false
		}
	}

	// 若栈不为空说明字符串末尾有多余的右括号
	return len(stack) == 0
}
