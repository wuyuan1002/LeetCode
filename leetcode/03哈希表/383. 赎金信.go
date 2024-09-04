package main

// 383. 赎金信

// 给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，
// 判断第一个字符串 ransom 能不能由第二个字符串 magazines 里面的字符构成。
// 如果可以构成，返回 true ；否则返回 false。
//
// (题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，
// 组成单词来表达意思。杂志字符串中的每个字符只能在赎金信字符串中使用一次。)

// canConstruct .
// leetcode 242. 有效的字母异位词
//
// 使用map存第二个字符串字符出现的次数
func canConstruct(ransomNote string, magazine string) bool {
	// 使用数组记录magazine每个字符出现的次数
	record := make([]int, 26)
	for _, c := range magazine {
		record[c-'a']++
	}

	// 遍历ransomNote判断, 能否由magazine的字符构成
	for _, c := range ransomNote {
		record[c-'a']--
		if record[c-'a'] < 0 {
			return false
		}
	}
	return true
}
