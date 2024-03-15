package main

// 567. 字符串的排列

// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
//
// 换句话说，s1 的排列之一是 s2 的 子串 。

// checkInclusion .
// 滑动窗口，判断窗口中各个字符出现的次数是否与s2中相同
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	// 构造两个数组用来统计窗口内两个字符串各个字符的出现次数
	var cnt1, cnt2 [26]int

	// 初始化窗口，窗口长度为s1的长度
	for i, c := range s1 {
		cnt1[c-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}

	// 不断向后移动窗口，更新窗口内的字符个数，并进行比较
	for i := len(s1); i < len(s2); i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-len(s1)]-'a']--

		if cnt1 == cnt2 {
			return true
		}
	}

	return false
}
