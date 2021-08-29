package main

// 3. 无重复字符的最长子串

func main() {

}

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
