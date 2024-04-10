package main

// 1768. 交替合并字符串

// 给你两个字符串 word1 和 word2 。请你从 word1 开始，通过交替添加字母来合并字符串。
// 如果一个字符串比另一个字符串长，就将多出来的字母追加到合并后字符串的末尾。
// 返回 合并后的字符串 。

// mergeAlternately .
// 使用一个指针指向当前要被合并的字符下标，分别将两字符串的当前下标字符加入到结果集，
// 若超出长度则不进行操作，直到指针将两个字符串长度都超过则循环结束
func mergeAlternately(word1 string, word2 string) string {
	result := make([]byte, 0, len(word1)+len(word2))

	// 不断向前移动指针，将两字符串的对应位置字符加入到结果集
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			result = append(result, word1[i])
		}
		if i < len(word2) {
			result = append(result, word2[i])
		}
	}

	return string(result)
}
