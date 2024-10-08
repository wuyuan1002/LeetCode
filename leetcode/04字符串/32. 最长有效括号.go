package main

// 32. 最长有效括号

// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

// longestValidParentheses .
//
// 遍历字符串的每个字符，使用两个变量存储左括号和右括号的出现次数，当左右括号出现次数相同时说明找到了一个有效的括号组合，更新有效括号的长度
// 当右括号比左括号多时，说明记录的括号已经不再合法（因为多了一个右括号是没有左括号与它匹配的），此时将左右括号的计数归零重新统计后续的字符
//
// 需要正序遍历一次，再倒序遍历一次，防止漏掉右括号个数始终比左括号个数多的情况，如"(()"
func longestValidParentheses(s string) int {
	maxLen := 0         // 最长有效括号
	left, right := 0, 0 // 左右括号出现次数

	for i := 0; i < len(s); i++ {
		// 更新左右括号出现次数
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if right == left {
			// 当右括号个数与左括号相同时说明当前构成的括号是有效的，更新有效括号的最大长度
			maxLen = max(maxLen, left+right)
		} else if right > left {
			// 当右括号个数比左括号多时，说明此时括号组合已经不合法了，将左右括号的计数归零重新统计后续的字符
			left, right = 0, 0
		}
	}

	// 倒序遍历一次字符串, 统计"(()"这种左括号始终大于右括号的情况
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		// 当左右括号个数一致时更新最长有效括号长度，当左括号个数大于右括号个数时将括号计数归零重新统计
		if left == right {
			maxLen = max(maxLen, left+right)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return maxLen
}
