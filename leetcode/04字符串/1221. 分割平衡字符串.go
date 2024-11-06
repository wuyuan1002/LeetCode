package main

// 1221. 分割平衡字符串

// 平衡字符串 中，'L' 和 'R' 字符的数量是相同的。
//
// 给你一个平衡字符串 s，请你将它分割成尽可能多的子字符串，并满足：
//
// 每个子字符串都是平衡字符串。
// 返回可以通过分割得到的平衡字符串的 最大数量 。

// balancedStringSplit .
//
// 因为原字符串不能拆分分割，只能按顺序从头到尾分割，所以直接从头到尾遍历字符串当L和R个数一致时便是找到了一个字串，
// 使用一个变量记录当前字串中L和R的个数差异，当差异为0时说明分割出了一个字串，然后继续遍历分割剩下的字串
func balancedStringSplit(s string) int {
	result, diff := 0, 0

	for _, c := range s {
		// 记录当前字串的L和R是否平衡
		if c == 'R' {
			diff++
		} else {
			diff--
		}

		// 当当前字串平衡时记录找到了一个平衡字串
		if diff == 0 {
			result++
		}
	}

	return result
}
