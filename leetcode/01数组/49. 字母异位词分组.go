package main

import "sort"

// 49. 字母异位词分组

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

// groupAnagrams .
// 字母异位词排序后为相等的字符串
// 先将所有字符串排序后存入map，最终同一个key所对应的就是一组异位词分组
func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)        // 总结果
	sortMap := make(map[string][]string) // 各个字符串排序，暂存异位词字符串（key为排序后的字符串）

	// 遍历每一个字符串，排序后存入map中
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortMap[string(s)] = append(sortMap[string(s)], str)
	}

	// map中的每一个key对应的value就是一组异位词分组，将他们添加到result中
	for _, group := range sortMap {
		result = append(result, group)
	}

	return result
}
