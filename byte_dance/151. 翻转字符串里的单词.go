package main

import (
	"fmt"
	"strings"
)

// 151. 翻转字符串里的单词

// 给你一个字符串 s ，逐个翻转字符串中的所有单词 。
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
// 请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。
//
// 说明：
// 输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
// 翻转后单词间应当仅用一个空格分隔。
// 翻转后的字符串中不应包含额外的空格。

func main() {
	fmt.Println(reverseWords(" the sky   is  blue"))
}

// offer 58
// 先翻转整个字符串，再翻转每一个单词, 或者先翻转每个单词，再翻转整个字符串
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

	// 先反转每个单词
	start := 0 // 当前单词第一个字母的下标
	for i := 1; i < len(sbyte); i++ {
		if sbyte[i] == ' ' {
			reverse(start, i-1)
			start = i + 1
		} else if i == len(sbyte)-1 { // 到了最后一个单词
			reverse(start, i)
		}
	}

	// 再反转整个字符串
	reverse(0, len(sbyte)-1)

	// 去掉单词间的多余空格
	str := strings.Builder{}
	for i := 0; i < len(sbyte); i++ {
		if sbyte[i] == ' ' && sbyte[i-1] == ' ' {
			continue
		}
		str.WriteByte(sbyte[i])
	}

	return str.String()
}
