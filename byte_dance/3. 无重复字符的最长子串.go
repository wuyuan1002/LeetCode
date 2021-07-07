package main

// 3. 无重复字符的最长子串

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

func main() {

}

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
