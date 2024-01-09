package main

// 438. 找到字符串中所有字母异位词

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
// 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

// findAnagrams .
// 滑动窗口+(数组或map)
// 使用两个map记录字符串和窗口中字符出现的次数，当窗口长度与字符串长度相同时，判断两个map中各个字符出现的次数是否一致，
// 若一致则说明此时窗口中的字串是字母异位词，将下标记入总结果
func findAnagrams(s string, p string) []int {
	result := make([]int, 0) // 总结果
	pcnt := make([]int, 26)  // 记录p串中字符出出现的次数
	scnt := make([]int, 26)  // 记录s串的窗口中字符出现的次数

	// 先将p串需要的字符个数入数组
	for _, c := range p {
		pcnt[c-'a']++
	}

	// 不断向右移动窗口，判断窗口内子串是否为字母异位词
	for i := 0; i < len(s); i++ {
		if i < len(p) {
			// 构建第一个窗口，字符直接入数组
			scnt[s[i]-'a']++
			continue
		}

		// 若窗口中的各个字符出现次数和目标字符串的各个字符出现次数一致，说明此时窗口中的子串是字母异位词
		if isEquals(scnt, pcnt) {
			result = append(result, i-len(p))
		}

		// 窗口右边界新进来的元素个数++，窗口左边出去的元素个数--
		scnt[s[i]-'a']++
		scnt[s[i-len(p)]-'a']--
	}

	// 统计最后一个窗口的子串
	if isEquals(scnt, pcnt) {
		result = append(result, len(s)-len(p))
	}

	return result
}

// isEquals 判断两个数组是否相等
func isEquals(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
