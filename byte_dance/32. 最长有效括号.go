package main

// 32. 最长有效括号

// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

// func main() {

// }

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	maxLen := 0

	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			right++
		} else {
			panic("遇到非法字符")
		}

		if left == right {
			maxLen = max(maxLen, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			right++
		} else {
			panic("遇到非法字符")
		}

		if left == right {
			maxLen = max(maxLen, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return maxLen
}
