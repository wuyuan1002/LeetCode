package main

import "strings"

// 151. 翻转字符串里的单词

func main() {

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

	// 反转每个单词
	start := 0
	for i := 0; i < len(sbyte); i++ {
		if sbyte[i] == ' ' {
			reverse(start, i-1)
			start = i + 1
		} else if i == len(sbyte)-1 {
			reverse(start, i)
		}
	}

	// 反转整个字符串
	reverse(0, len(sbyte)-1)

	// 去掉中间的多余空格
	sb := strings.Builder{}
	for i := 0; i < len(sbyte); i++ {
		if sbyte[i] == ' ' && sbyte[i-1] == ' ' {
			continue
		}
		sb.WriteByte(sbyte[i])
	}

	return sb.String()
}
