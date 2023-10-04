package main

// 28. 找出字符串中第一个匹配项的下标

// 给你两个字符串 haystack 和 needle ，
// 请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
// 如果 needle 不是 haystack 的一部分，则返回  -1 。

// strStr .
// 同 459. 重复的子字符串
// 1. KMP前缀表解法 -- 使用前缀数组记录子串指针匹配失败时子串指针最多可以向前移动几个字符
// 2. 遍历第一个字符串, 每遍历一个字符接着遍历第二个字符串进行比较
func strStr(haystack string, needle string) int {
	index := -1 // 结果

	// 遍历第一个字符串的每一个字符
	for i := 0; i < len(haystack); i++ {
		// 定义两个指针同时向后遍历两个字符串
		j, k := i, 0
		for ; j < len(haystack) && k < len(needle) && haystack[j] == needle[k]; j, k = j+1, k+1 {
		}
		// 若遍历到第二个字符串末尾，说明在第一个字符串中找到了第二个字符串 -- i为子串起始下表
		if k == len(needle) {
			index = i
			break
		}
	}

	return index
}
