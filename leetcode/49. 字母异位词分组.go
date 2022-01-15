package main

import (
	"sort"
)

// 49. 字母异位词分组

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
