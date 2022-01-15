package main

// 3. 无重复字符的最长子串

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

// func main() {

// }

// 滑动窗口，同时使用map存储字符出现过的下标
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	sbyte := []byte(s)
	maxLen := 0                // 最大子串长度
	hash := make(map[byte]int) // 存储遍历到的字符和下标
	l, r := 0, 0               // 左右指针

	for ; r < len(sbyte); r++ {
		if index, ok := hash[sbyte[r]]; ok {
			// 若字符之前出现过，更新窗口左边界
			l = max(index+1, l)
		}
		// 将当前右边界字符加入到map中，标记已经在子串中出现过
		hash[sbyte[r]] = r

		// 更新子串长度
		maxLen = max(maxLen, r-l+1)
	}

	return maxLen
}

// 第二次做
// 滑动窗口，同时使用map存储字符出现过的下标
func lengthOfLongestSubstring1(s string) int {
	if s == "" {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	hash := make(map[byte]int) // 存储字符出现过的下标
	maxLen := 0                // 最长子串长度
	l := 0                     // 左指针
	for i, c := range s {
		if index, ok := hash[byte(c)]; ok {
			// 更新窗口左边界
			l = max(l, index+1)
		}
		// 更新窗口有边界
		hash[byte(c)] = i

		// 更新最大长度
		maxLen = max(maxLen, i-l+1)
	}

	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	if s == "" {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

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
