package main

// 5. 最长回文子串

// 给你一个字符串 s，找到 s 中最长的回文子串。

// longestPalindrome1 .
// 同 leetcode 647、516 -- 回文子串需要连续，回文子序列不需要连续
//
// 1. 中心拓展法: 遍历数组，先左右寻找与当前字符相等的字符，之后左右同时扩散，直到左右指针的值不相等
func longestPalindrome1(s string) string {
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

// longestPalindrome2 .
// 同 leetcode 132. 分割回文串 II
//
// 2. 动态规划: 使用dp数组缓存子串是否为回文串 -- dp[i][j]表示s[i, j]子串是否为回文串
func longestPalindrome2(s string) string {
	// 构造dp数组，用于缓存中间子串是否为回文串 -- dp[i][j]表示s[i, j]子串是否为回文串
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	// 最长回文串的起点、最长回文串的终点
	maxStart, maxEnd := 0, 0

	// 开始dp，遍历左右指针，计算每个子串s[l,r]是否为回文串，同时更新最长回文串的起止下标
	for r := 0; r < len(s); r++ {
		for l := 0; l <= r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				// 标记当前左右指针间的子串s[l,r]为回文串
				dp[l][r] = true

				// 若当前子串长度比记录的最长子串更长，则更新当前子串为最长回文串
				if r-l+1 > maxEnd-maxStart+1 {
					maxStart, maxEnd = l, r
				}
			}
		}
	}

	return s[maxStart : maxEnd+1]
}

// longestPalindrome3 .
// 同 leetcode 1143. 最长公共子序列
//
// 3. 回文串正反都一样，所以可以求字符串s和将s倒置后的字符串求最长公共子串即可
func longestPalindrome3(s string) string {
	longestCommonSubsequence(s, s)
	return ""
}
