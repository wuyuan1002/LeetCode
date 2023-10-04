package main

import "strings"

// 459. 重复的子字符串

// 给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。
// 给定的字符串只含有小写英文字母，并且长度不超过10000。

// repeatedSubstringPattern .
// 同 28. 找出字符串中第一个匹配项的下标
// 1. KMP前缀表解法
// 2. 移动匹配：将两个 s 连在一起，并移除第一个和最后一个字符。如果 s 是该字符串的子串，那么 s 就满足题目要求
// 如果 s 中没有循环节，那么 ss 中必然有且只有两个 s，
// 此时从 ss[1] 处开始寻找 s ，必然只能找到第二个，所以此时返回值为 s.size()。
// 如果s中有循环节，查找的返回值将是第一个s中的一个循环节下标，必然小于 s.size()
func repeatedSubstringPattern(s string) bool {
	// 移动匹配
	// 构造一个新的字符串由两个s拼接而成，然后掐头去尾
	// 判断原字符串s是否为新字符串的字串
	str := (s + s)[1 : 2*len(s)-1]
	return strings.Contains(str, s)
}
