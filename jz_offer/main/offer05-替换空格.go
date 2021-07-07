package main

import (
	"fmt"
	"strings"
)

// 剑指 Offer 05. 替换空格
// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

func main() {
	fmt.Printf("%s", replaceSpace2("alkcn ndc s "))
}

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

// 第二次做
func replaceSpace2(s string) string {
	if s == "" {
		return ""
	}

	l := 0 // 字符串中空格的个数
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			l++
		}
	}

	aa := make([]byte, l*2)
	// copy([]byte(s), sb)
	sb := append([]byte(s), aa...)

	i := len(sb) - 1 // 指向数组的末尾
	j := len(s) - 1  // 指向数组中最后一个字母
	for i != j && i >= 0 && j >= 0 {
		if sb[j] == ' ' {
			sb[i] = '0'
			i--
			sb[i] = '2'
			i--
			sb[i] = '%'
		} else {
			sb[i] = sb[j]
		}
		i--
		j--
	}
	return string(sb)
}
