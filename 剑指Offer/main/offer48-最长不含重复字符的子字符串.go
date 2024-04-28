package main

// 剑指 Offer 48. 最长不含重复字符的子字符串

// 请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

// func main() {
// 	// fmt.Println(lengthOfLongestSubstring("pwwkew"))
// 	// fmt.Println(lengthOfLongestSubstring("abba"))
// 	fmt.Println(lengthOfLongestSubstring1("abcabcbb"))
// }

// 同 hot100 3
// 使用滑动窗口解决
func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 存储 left ~ right 中间的字符,以及其出现的下标
	set := make(map[byte]int)

	left := 0
	right := 1
	maxLen := 1
	set[s[left]] = 0

	for ; right < len(s); right++ {
		if _, ok := set[s[right]]; ok {
			left = max(set[s[right]]+1, left) // 把左指针移到之前下标的下一位; left不能被重置到比当前left前面的位置，可以debug案例 abba
		}
		set[s[right]] = right // 若窗口中没有，则存进去，若存在，则更新之前存在字母的下标

		maxLen = max(right-left+1, maxLen)
	}
	return maxLen
}

// 第二次做
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	l := 0      // 左指针
	r := 1      // 右指针
	maxlen := 1 // 最大长度

	set := make(map[byte]int) // 存储窗口中字符和下标
	set[s[l]] = 0

	for ; r < len(s); r++ {
		if _, ok := set[s[r]]; ok { // 若当前字符出现过 -- 重置左指针的位置
			l = max(set[s[r]]+1, l) // 把左指针移到当前字符上一次出现的下一位;l不能被重置到比当前l前面的位置，可以debug案例 abba
		}
		set[s[r]] = r // 若之前出现过则更新当前字符出现的下标，若之前没有出现过则新加入到集合中

		maxlen = max(r-l+1, maxlen) // 重新计算最大长度
	}

	return maxlen
}
