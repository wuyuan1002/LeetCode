package main

import "strings"

// 459. 重复的子字符串

// 给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。
// 给定的字符串只含有小写英文字母，并且长度不超过10000。

// func main() {

// }

func repeatedSubstringPattern(s string) bool {
	str := (s + s)[1 : 2*len(s)-1]
	return strings.Contains(str, s)
}
