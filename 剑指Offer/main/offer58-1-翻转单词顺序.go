package main

import (
	"fmt"
	"strings"
)

// 剑指 Offer 58 - I. 翻转单词顺序

// 输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
// 为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。

// func main() {
// 	// fmt.Println(reverseWords(" I am a student. "))
// 	fmt.Println(reverseWords2("a good  example"))
// }

// 先翻转整个字符串，再翻转每一个单词, 或者先翻转每个单词，再翻转整个字符串
func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}

	s = strings.TrimSpace(s) // 去掉两边空格

	str := []byte(s)

	reverse(0, len(str)-1, str) // 先翻转整个字符串

	i, j := 0, 0
	for j < len(str) {
		if str[j] == ' ' { // 根据空格，翻转每一个单词
			reverse(i, j-1, str)
			i = j + 1
		} else if j == len(str)-1 { // 若到了最后一个字符，直接翻转这个单词
			reverse(i, j, str)
		}
		j++
	}
	return string(str)
}

// 左闭右闭
func reverse(i int, j int, str []byte) {
	if i < 0 || j > len(str)-1 {
		return
	}
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
}

func reverseWords1(s string) string {
	strSlice := strings.Fields(s)
	fmt.Println(strSlice)
	left, right := 0, len(strSlice)-1
	for left < right {
		strSlice[left], strSlice[right] = strSlice[right], strSlice[left]
		left++
		right--
	}

	return strings.Join(strSlice, " ")
}

func reverseWords2(s string) string {
	var strs []string
	var fast int
	for fast < len(s) {
		var str string
		// 找到第一个不是空格的
		for fast < len(s) && s[fast] == ' ' {
			fast++
		}
		// 找到下一个空格
		for fast < len(s) && s[fast] != ' ' {
			str += string(s[fast])
			fast++
		}
		if len(str) > 0 {
			strs = append(strs, str)
		}
	}
	for i := 0; i < len(strs)/2; i++ {
		strs[i], strs[len(strs)-i-1] = strs[len(strs)-i-1], strs[i]
	}
	return strings.Join(strs, " ")
}
