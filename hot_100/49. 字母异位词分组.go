package main

import (
	"sort"
)

// 49. 字母异位词分组

// 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

func main() {

}

// 1. 对每个字符串排序后，字母异位词就会相等，直接比较即可
// 2. 挨个遍历每个字符串，将字符串对应字符和出现的次数记录在map中，再遍历剩下的字符串比较和当前字符串是否构成字母异位词

// 1.
func groupAnagrams(strs []string) [][]string {
	if strs == nil || len(strs) == 0 {
		return nil
	}

	mp := make(map[string][]string) // key: 排序后的字符串，value: 每个字符串
	// 遍历每一个字符串
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}

	result := make([][]string, 0, len(mp))
	for _, res := range mp {
		result = append(result, res)
	}

	return result
}
