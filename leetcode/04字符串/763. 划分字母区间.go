package main

// 763. 划分字母区间

// 给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。
//
// 注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。
//
// 返回一个表示每个字符串片段的长度的列表。

// partitionLabels .
// - 首先遍历字符串，记录字符串中每个字符出现的最远下标
// - 再次遍历字符串，若当前字符出现的最远下标比当前区间的结束位置更远，则更新当前区间的结束位置
// - 若当前字符的下标已经到达当前区间的结束位置，则说明当前区间可以被切割了
func partitionLabels(s string) []int {
	result := make([]int, 0)

	// 使用数组记录字符串中每个字符出现的最远下标
	lastPos := [26]int{}
	for i, c := range s {
		lastPos[c-'a'] = i
	}

	// 遍历字符串的每个字符，进行区间切割
	// start表示当前区间的起始下标，end表示当前区间的结束下标
	start, end := 0, 0
	for i, c := range s {
		// 若当前字符出现的最远下标比目前记录的当前区间结束位置更远，则更新当前区间的结束位置
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}

		// 若当前遍历到的字符下标已经到达当前区间的结束位置，说明当前区间可被切割 -- 切割当前区间并开始下一个区间的切割
		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}

	return result
}
