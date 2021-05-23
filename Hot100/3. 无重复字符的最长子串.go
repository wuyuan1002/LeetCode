package main

import (
	"fmt"
)

// 给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	// fmt.Println(lengthOfLongestSubstring("abba"))
	// fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

// 同offer 48
// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sbyte := []byte(s)
	maxlen := 0                // 记录最大长度
	hash := make(map[byte]int) // 存储遍历到的字符和下标
	l, r := 0, 0               // 左右指针
	for ; r < len(sbyte); r++ {
		if index, ok := hash[sbyte[r]]; ok { // 若该字符之前出现过
			l = max(l, index+1)
		}
		hash[sbyte[r]] = r // 标记或更新当前字符的下标

		maxlen = max(maxlen, r-l+1)
	}

	return maxlen
}
