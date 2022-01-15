package main

// 438. 找到字符串中所有字母异位词

// 给定一个字符串s和一个非空字符串p，找到s中所有是p的字母异位词的子串，返回这些子串的起始索引。
// 字符串只包含小写英文字母，并且字符串s和 p的长度都不超过 20100。
//
// 说明：
// 字母异位词指字母相同，但排列不同的字符串。
// 不考虑答案输出的顺序。

// func main() {
// 	fmt.Println(findAnagrams("cbaebabacd", "abc"))
// 	// fmt.Println(findAnagrams("abab", "ab"))
// }

// 滑动窗口+(数组或map)
func findAnagrams(s string, p string) []int {
	if s == "" || p == "" || len(s) < len(p) {
		return nil
	}

	result := make([]int, 0) // 总结果
	pcnt := make([]int, 26)  // 记录p串中字符出出现的次数
	scnt := make([]int, 26)  // 记录s串的窗口中字符出现的次数

	isEquals := func() bool { // 判断两个数组是否相等
		for i := 0; i < 26; i++ {
			if pcnt[i] != scnt[i] {
				return false
			}
		}
		return true
	}

	// 先将p串需要的字符个数入数组
	for _, c := range p {
		pcnt[c-'a']++
	}

	// 窗口遍历s串
	for i := 0; i < len(s); i++ {
		if i < len(p) {
			// 构建第一个窗口，字符直接入数组
			scnt[s[i]-'a']++
			continue
		}

		if isEquals() {
			result = append(result, i-len(p))
		}

		scnt[s[i]-'a']++        // 窗口右边界新进来的元素个数++
		scnt[s[i-len(p)]-'a']-- // 窗口左边出去的元素个数--
	}

	if isEquals() {
		result = append(result, len(s)-len(p))
	}

	return result
}
