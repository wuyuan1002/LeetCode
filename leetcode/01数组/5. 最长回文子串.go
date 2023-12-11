package main

// 5. 最长回文子串

// 给你一个字符串 s，找到 s 中最长的回文子串。

// longestPalindrome .
// 同 leetcode 647、516 -- 回文子串需要连续，回文子序列不需要连续
// 1.Manacher算法 见Hot100 647
// 2.中心拓展法: 遍历数组，先 左右寻找与当前字符相等的字符，之后左右同时扩散，直到左右指针的值不相等
func longestPalindrome(s string) string {
	// 总结果 -- 最长回文串
	result := ""

	for i := 0; i < len(s); i++ {
		l, r := i-1, i+1

		// 先向左和向右遍历与当前字符相等的字符
		for l >= 0 && s[l] == s[i] {
			l--
		}
		for r <= len(s)-1 && s[r] == s[i] {
			r++
		}

		// 之后左右指针同时向两边扩散，直至左右指针的值不相等
		for l >= 0 && r <= len(s)-1 && s[l] == s[r] {
			l--
			r++
		}

		// 更新最长回文串
		if r-l-1 > len(result) {
			result = s[l+1 : r]
		}
	}

	return result
}
