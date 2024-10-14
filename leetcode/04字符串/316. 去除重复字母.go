package main

// 316. 去除重复字母

// 给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。
// 需保证 返回结果的字典序 最小（要求不能打乱其他字符的相对位置）。

// removeDuplicateLetters .
//
// 遍历一次字符串，统计每个字符出现的次数，用于记录下一步遍历生成结果时后续剩余的字符数量
// 再次遍历字符串，生成结果，不断将当前字母与结果中的前面字符比较字典序，若小于前面字符且前面字符剩余数量大于0则使用当前字符替换掉前面字符
func removeDuplicateLetters(s string) string {
	// 遍历一遍字符串，统计每个字符出现的剩余次数
	count := make(map[byte]int)
	for _, c := range []byte(s) {
		count[c]++
	}

	// 总结果、用于记录字符是否已在总结果中
	result, isInReault := make([]byte, 0), make(map[byte]bool)

	// 再次遍历字符串，计算最终结果
	for _, c := range []byte(s) {
		// 递减当前字符剩余的出现次数
		count[c]--

		// 若当前字符在结果中已经出现过则直接跳过
		if isInReault[c] {
			continue
		}

		// 从后向前去除总结果中所有大于当前字符且还有剩余的字符
		// 即将字典序小的字符尽量安排到总结果的前面，保证总结果的字典序最小
		for len(result) > 0 && c < result[len(result)-1] && count[result[len(result)-1]] > 0 {
			// 更新标记末尾字符在总结果中未出现，并将其从总结果中去除
			isInReault[result[len(result)-1]] = false
			result = result[:len(result)-1]
		}

		// 将当前字符加入到总结果末尾，并标记当前字符在总结果中出现过
		result = append(result, c)
		isInReault[c] = true
	}

	return string(result)
}
