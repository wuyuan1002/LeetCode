package main

import (
	"sort"
)

// 49. 字母异位词分组

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

// func main() {

// }

// 字母异位词排序后为相等的字符串
func groupAnagrams(strs []string) [][]string {
	if strs == nil || len(strs) == 0 {
		return nil
	}

	mapp := make(map[string][]string, 0)
	result := make([][]string, 0)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(a, b int) bool { return s[a] < s[b] })
		mapp[string(s)] = append(mapp[string(s)], str)
	}

	for _, v := range mapp {
		result = append(result, v)
	}
	return result
}

// 第二次做
func groupAnagrams1(strs []string) [][]string {
	if strs == nil || len(strs) == 0 {
		return nil
	}

	result := make([][]string, 0)        // 总结果
	mapp := make(map[string][]string, 0) // 存放字母异位词分组

	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(a, b int) bool { return s[a] < s[b] })
		bystr := string(s)
		mapp[bystr] = append(mapp[bystr], str)
	}

	for _, v := range mapp {
		result = append(result, v)
	}

	return result
}
