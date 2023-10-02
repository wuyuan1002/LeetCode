package main

// 3. 无重复字符的最长子串

// 给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。

// func main() {
// 	fmt.Println(lengthOfLongestSubstring("pwwkew"))
// 	// fmt.Println(lengthOfLongestSubstring("abba"))
// 	// fmt.Println(lengthOfLongestSubstring("abcabcbb"))
// }

// 滑动窗口，同时使用map存储字符出现过的下标
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	hash := make(map[byte]int, 0)
	maxLen := 0
	l := 0
	for i, c := range s {
		if index, ok := hash[byte(c)]; ok {
			l = max(l, index+1)
		}
		hash[byte(c)] = i
		maxLen = max(maxLen, i-l+1)
	}
	return maxLen
}

func lengthOfLongestSubstring1(s string) int {
	hash := make(map[byte]int)
	maxLen := 0
	l := 0
	for i, c := range s {
		if index, ok := hash[byte(c)]; ok {
			l = max(l, index+1)
		}
		hash[byte(c)] = i
		maxLen = max(maxLen, i-l+1)
	}
	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	hash := make(map[byte]int)
	maxLen := 0
	l := 0
	for i, c := range s {
		if index, ok := hash[byte(c)]; ok {
			// 若该字符已经存在，将左边界向前移动，不能直接移动到index+1，原因：abba
			l = max(l, index+1)
		}
		// 将最新位置存入map
		hash[byte(c)] = i
		// 计算最大长度
		maxLen = max(maxLen, i-l+1)
	}
	return maxLen
}
