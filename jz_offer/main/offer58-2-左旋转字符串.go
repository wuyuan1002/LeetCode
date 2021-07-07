package main

import (
	"strings"
)

// 剑指 Offer 58 - II. 左旋转字符串

// 字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
// 请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，
// 该函数将返回左旋转两位得到的结果"cdefgab"。
func main() {

}

// 先翻转整个字符串，再翻转每一部分, 或者先翻转每个部分，再翻转整个字符串
func reverseLeftWords(s string, n int) string {
	if len(s) == 0 || n >= len(s) {
		return s
	}

	s = strings.TrimSpace(s)

	str := []byte(s)
	reverse(0, n-1, str)      // 翻转前n部分
	reverse(n, len(s)-1, str) // 翻转n后面的部分
	reverse(0, len(s)-1, str) // 翻转整个字符串
	return string(str)
}
