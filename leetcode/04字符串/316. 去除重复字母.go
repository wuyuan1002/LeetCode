package main

// 316. 去除重复字母

// 给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。
// 需保证 返回结果的字典序 最小（要求不能打乱其他字符的相对位置）。

// removeDuplicateLetters .
// - 遍历一遍字符串，统计每个字符出现的次数 -- 用于记录下一步遍历生成结果时后续剩余的字符数量
// - 再次遍历字符串，生成结果，不断将当前字母与结果中的前面字符比较字典序，若小于前面字符且前面字符剩余数量大于0则使用当前字符替换掉前面字符
func removeDuplicateLetters(s string) string {
	// 遍历一遍字符串，统计每个字符出现的剩余次数
	count := make([]int, 26)
	for _, c := range []byte(s) {
		count[c-'a']++
	}

	// 中间结果（存的是字符减去'a'的ascii码）、用于记录字符是否已在总结果中
	res, isInReault := make([]byte, 0), make([]bool, 26)

	// 再次遍历字符串，计算最终结果
	for _, c := range []byte(s) {
		// 递减当前字符剩余的出现次数
		count[c-'a']--

		// 若当前字符在结果中已经出现过则直接跳过
		if isInReault[c-'a'] {
			continue
		}

		// 若当前字符小于结果的末尾字符且末尾字符还有剩余，则将末尾字符去除 -- 从后向前去除总结果中所有大于当前字符且还有剩余的字符
		for len(res) > 0 && c-'a' < res[len(res)-1] && count[res[len(res)-1]] > 0 {
			// 将末尾字符从总结果中去除并更新标记在总结果中未出现
			isInReault[res[len(res)-1]] = false
			res = res[:len(res)-1]
		}

		// 将当前字符加入到总结果末尾并标记当前字符在总结果中出现过
		res = append(res, c-'a')
		isInReault[c-'a'] = true
	}

	// 最终结果 -- 将每个字符的ascii码加'a'
	result := make([]rune, len(res))
	for i, c := range res {
		result[i] = rune(c + 'a')
	}

	return string(result)
}
