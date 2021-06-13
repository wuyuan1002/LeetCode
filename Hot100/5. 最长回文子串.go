package main

import (
	"fmt"
)

// 5. 最长回文子串

// 给你一个字符串 s，找到 s 中最长的回文子串。

func main() {
	// fmt.Println([]int{1,2,3,4,5}[3:5])
	fmt.Println(longestPalindrome("babad"))
}

// 遍历数组，先 左右寻找与当前字符相等的字符，之后左右同时扩散，知道做右指针的值不相等
func longestPalindrome(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}

	minl, maxr := 0, 0 // 最长回文子串的起始结束下标
	maxlen := 1        // 最长回文串长度
	nowlen := 1        // 当前回文串长度

	l, r := 0, 0 // 左右指针

	sbyte := []byte(s)
	for i := 0; i < len(sbyte)-1; i++ {
		l, r = i-1, i+1
		for l > 0 && sbyte[l] == sbyte[i] {
			nowlen++
			l--
		}
		for r < len(sbyte) && sbyte[r] == sbyte[i] {
			nowlen++
			r++
		}
		for l >= 0 && r < len(sbyte) && sbyte[l] == sbyte[r] {
			nowlen += 2
			l--
			r++
		}
		if nowlen > maxlen {
			maxlen = nowlen
			minl, maxr = l+1, r-1
		}
		nowlen = 1
	}
	return string(sbyte[minl : maxr+1])
}
