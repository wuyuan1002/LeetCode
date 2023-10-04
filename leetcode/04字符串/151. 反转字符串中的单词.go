package main

// 151. 反转字符串中的单词

// 给你一个字符串 s ，逐个翻转字符串中的所有单词 。
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
// 请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。
//
// 说明：
// 输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
// 翻转后单词间应当仅用一个空格分隔。
// 翻转后的字符串中不应包含额外的空格。

// reverseWords .
// offer 58、leetcode 151、189、541
// 先反转整个字符串，再反转每一个单词, 或者先反转每个单词，再反转整个字符串
func reverseWords(s string) string {

	// 去除前面、中间、后面存在的多余空格
	// s = strings.TrimSpace(s)
	sb := removeExtraSpaces([]byte(s))
	if len(sb) == 0 {
		return ""
	}

	// 反转整个字符串
	reverse(sb)

	// 反转每一个单词
	for start, i := 0, 0; i <= len(sb); i++ {
		// i到达末尾或指向空格时, 开始反转前一个单词
		// 没到达末尾时, i指向的是前一个单词后的空格,
		// 到达末尾时, i指向的是最后一个字母的后一个位置, 所以i-1始终是前一个单词的最后一个字母
		if i == len(sb) || sb[i] == ' ' {
			// 反转前一个单词
			reverse(sb[start:i])
			// 将单词开始位置指向下一个单词的首字母, i此时指向的是空格, i+1指向的便是下一个单词的首字母
			start = i + 1
		}
	}

	return string(sb)
}

// removeExtraSpaces 移除所有空格并在相邻单词之间添加空格 -- 双指针
// 同 27.移除元素 -- 只是在移除元素的过程中保留手动保留一个空格
func removeExtraSpaces(s []byte) []byte {
	if s == nil || len(s) == 0 {
		return s
	}

	i, j := 0, 0
	for j < len(s) {
		if s[j] != ' ' {
			// 特殊逻辑, 在遍历到后一个单词的时候, 为前一个单词末尾增加空格
			// 第一个单词前面不用加, i > 0 说明不是第一个单词
			if i > 0 && j > 0 && s[j-1] == ' ' {
				s[i] = ' '
				i++
			}

			s[i] = s[j]
			i++
		}
		j++
	}
	return s[:i]
}
