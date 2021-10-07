package main

// 434. 字符串中的单词数

// 统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
//
// 请注意，你可以假定字符串里不包括任何不可打印的字符。

func main() {

}

// 统计前一个是空格且当前位置不是空格的字符个数 -- 统计单词的第一个字母的个数
func countSegments(s string) int {
	if s == "" {
		return 0
	}

	num := 0
	for i := range s {
		if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
			num++
		}
	}
	return num
}
