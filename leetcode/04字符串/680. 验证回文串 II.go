package main

// 680. 验证回文串 II

// 给你一个字符串 s，最多 可以从中删除一个字符。
//
// 请你判断 s 是否能成为回文字符串：如果能，返回 true ；否则，返回 false 。

// validPalindrome .
// 左右指针分别从两边遍历比较，若出现两指针指向字符不一致则说明需要删除左指针或右指针指向的字符来使其成为回文串，
// 此时分别验证移除左指针和右指针指向的字符，看其是否能成为回文串
func validPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] == s[r] {
			continue
		}

		// 若左右指针指向的字符不一致了，则直接查看删掉左指针或右指针指向的字符能否构成回文串
		return isPalindrome680(s[l+1:r+1]) || isPalindrome680(s[l:r])
	}
	return true
}

// isPalindrome680 验证字符串是否为回文串
func isPalindrome680(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}
