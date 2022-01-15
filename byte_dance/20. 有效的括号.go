package main

// 20. 有效的括号

// 给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。

// func main() {

// }

var buckets = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}

	stack := make([]byte, 0)
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, byte(c))
		} else if len(stack) > 0 && (c == ')' || c == ']' || c == '}') {
			if byte(c) == buckets[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			// 遇到非法字符
			return false
		}
	}

	return len(stack) == 0
}
