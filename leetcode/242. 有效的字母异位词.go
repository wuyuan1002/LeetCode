package main

// 242. 有效的字母异位词

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。

func main() {

}

// 1. 排序两个字符串，然后比较是否相等
// 2. 统计两个字符串中字符的个数，所有字符出现次数相等，则是字母异位词
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	record := make([]int, 26) // 记录每个字符出现的次数
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
