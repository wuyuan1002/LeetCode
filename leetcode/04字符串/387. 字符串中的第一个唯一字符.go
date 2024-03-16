package main

// 387. 字符串中的第一个唯一字符

// 给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。

// firstUniqChar .
func firstUniqChar(s string) int {
	// 第一次遍历进行字符个数统计
	count := [26]int{}
	for _, c := range s {
		count[c-'a']++
	}

	// 第二次遍历返回第一个计数为1的字符下标
	for i, c := range s {
		if count[c-'a'] == 1 {
			return i
		}
	}
	return -1
}
