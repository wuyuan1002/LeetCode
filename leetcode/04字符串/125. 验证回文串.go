package main

import "strings"

// 125. 验证回文串

// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，
// 短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
//
// 字母和数字都属于字母数字字符。
//
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。

// isPalindrome .
//
// 使用双指针对字符串中的字母和数字进行比对，若遇到非字母数字字符则将指针向前移动至字母数字字符上
func isPalindrome(s string) bool {
	// 先将字符串所有字符转为小写
	s = strings.ToLower(s)

	l, r := 0, len(s)-1
	for l < r {
		// 若左右指针指向字符不是字母或数字，则将其移动到字母数字上
		for !isalnum(s[l]) && l < r {
			l++
		}
		for !isalnum(s[r]) && l < r {
			r--
		}

		if l < r {
			// 若左右指针指向字符不相同说明不是回文串
			if s[l] != s[r] {
				return false
			}

			// 将左右指针向前移动一位
			l++
			r--
		}
	}

	return true
}

// isalnum 判断给定字符是否为字母数字
func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
