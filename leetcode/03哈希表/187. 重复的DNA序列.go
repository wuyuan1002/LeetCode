package main

// 187. 重复的DNA序列

// DNA序列 由一系列核苷酸组成，缩写为 'A', 'C', 'G' 和 'T'.。
//
// 例如，"ACGAATTCCG" 是一个 DNA序列 。
// 在研究 DNA 时，识别 DNA 中的重复序列非常有用。
//
// 给定一个表示 DNA序列 的字符串 s ，返回所有在 DNA 分子中出现不止一次的 长度为 10 的序列(子字符串)。
// 你可以按 任意顺序 返回答案。

// findRepeatedDnaSequences .
//
// 滑动窗口
// 构造一个长度为10的窗口，不断向后滑动窗口，并将窗口所构成的字符串记录在哈希表中统计其出现的次数，
// 当发现一个窗口子串出现不止一次时将其计入结果集
func findRepeatedDnaSequences(s string) []string {
	result := make([]string, 0)
	count := make(map[string]int)

	for i := 0; i <= len(s)-10; i++ {
		sub := s[i : i+10]

		// 递增当前子串出现的次数
		count[sub]++

		// 若当前子串出现不止一次则记录到总结果 -- 此处只在出现第二次时记录到总结果，防止出现多次时被重复添加
		if count[sub] == 2 {
			result = append(result, sub)
		}
	}

	return result
}
