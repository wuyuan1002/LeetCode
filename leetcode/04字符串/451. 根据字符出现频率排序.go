package main

import "sort"

// 451. 根据字符出现频率排序

// 给定一个字符串 s ，根据字符出现的 频率 对其进行 降序排序 。一个字符出现的 频率 是它出现在字符串中的次数。
//
// 返回 已排序的字符串 。如果有多个答案，返回其中任何一个。

// frequencySort .
// 同 leetcode 1387. 将整数按权重排序
//
// 先遍历数组求出每个字符出现的次数，然后对数组按照字符出现的频率进行排序
func frequencySort(s string) string {
	bs := []byte(s)

	// 使用map记录各字符出现的次数
	count := make(map[byte]int)
	for _, c := range bs {
		count[c]++
	}

	// 对数组进行排序 -- 若出现的次数不同则按照出现的次数排序，若出现的次数相同则按照字符本身大小进行排序
	sort.Slice(bs, func(i, j int) bool {
		if count[bs[i]] != count[bs[j]] {
			// 若出现次数不同则直接按照次数进行排序 -- 决定出现次数不同的字符外部整体顺序
			return count[bs[i]] > count[bs[j]]
		} else {
			// 若两字符出现的次数相同，则按照字符本身大小进行排序 -- 决定相同出现次数字符间的内部顺序，确保出现次数相同的字符可以相同字符挨在一起 -- 如loveleetcode排序后是eeeeoollvtdc，而不是eeeelolovtcd
			return bs[i] > bs[j]
		}
	})

	return string(bs)
}
