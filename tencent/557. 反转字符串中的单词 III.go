package main

import (
	"strings"
)

// 557. 反转字符串中的单词 III

// 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

// func main() {
// 	fmt.Println(reverseWords(" the sky   is  blue"))
// }

// byte_dance 151 翻转字符串里的单词
func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	sbyte := []byte(s)
	reverse := func(l, r int) {
		for l < r {
			sbyte[l], sbyte[r] = sbyte[r], sbyte[l]
			l++
			r--
		}
	}

	// 反转每一个单词
	start := 0 // 当前单词第一个字母的下标
	for i := 0; i < len(sbyte); i++ {
		if sbyte[i] == ' ' {
			reverse(start, i-1)
			start = i + 1
		} else if i == len(sbyte)-1 { // 到了最后一个单词
			reverse(start, i)
		}
	}

	// 去掉字符串中多余的空格
	str := strings.Builder{}
	for i := 0; i < len(sbyte); i++ {
		if sbyte[i] == ' ' && sbyte[i-1] == ' ' {
			continue
		}
		str.WriteByte(sbyte[i])
	}

	return str.String()
}
