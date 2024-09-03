package main

// 242. 有效的字母异位词

// isAnagram .
// leetcode 383. 赎金信
//
// 1. 排序两个字符串，然后比较是否相等
// 2. 统计两个字符串中字符的个数，所有字符出现次数相等，则是字母异位词（可以使用数组也可以使用map进行记录）
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 使用数组记录a-z出现的次数
	record := make([]int, 26)

	for _, c := range s {
		record[c-'a']++
	}
	for _, c := range t {
		record[c-'a']--
	}

	for _, n := range record {
		if n != 0 {
			return false
		}
	}

	return true
}
