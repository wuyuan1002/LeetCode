package main

import (
	"fmt"
)

// 5. 最长回文子串

// 给你一个字符串 s，找到 s 中最长的回文子串。

func main() {
	fmt.Println(longestPalindrome("asdsa"))
}

// 中心扩散法
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	res := ""       // 最长回文串
	maxLen := 0     // 最大长度
	currentLen := 1 // 当前回文串长度
	for i, c := range s {
		l, r := i-1, i+1
		for l >= 0 && s[l] == byte(c) {
			currentLen++
			l--
		}
		for r <= len(s)-1 && s[r] == byte(c) {
			currentLen++
			r++
		}

		for l >= 0 && r <= len(s)-1 && s[l] == s[r] {
			currentLen += 2
			l--
			r++
		}

		if currentLen > maxLen {
			maxLen = currentLen
			res = s[l+1 : r]
		}

		currentLen = 1
	}

	return res
}
