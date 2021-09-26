package main

// 524. 通过删除字母匹配到字典里最长单词

// 给你一个字符串 s 和一个字符串数组 dictionary 作为字典，找出并返回字典中最长的字符串，
// 该字符串可以通过删除 s 中的某些字符得到。
//
// 如果答案不止一个，返回长度最长且字典序最小的字符串。如果答案不存在，则返回空字符串。

func main() {
	findLongestWord("abcb", []string{"aplpllre"})
}

// 类似1143
// 1. 双指针遍历每一个元素和给定字符串，返回匹配的最长串
// 2. 由于双指针有大量的时间用于在 s 中找到下一个匹配字符，
// 因此可以使用动态规划对于 s 的每一个i位置，保存从该位置开始往后每一个字符第一次出现的位置
func findLongestWord(s string, dictionary []string) string {
	if dictionary == nil || len(dictionary) == 0 {
		return ""
	}

	dp := make([][26]int, len(s)+1) // dp[i][j]表示s中从位置i开始往后字符j第一次出现的位置, 如果dp[i][j] == len(s)，则表示从位置i开始往后不存在字符j
	for i := range dp[len(s)] {
		dp[len(s)][i] = len(s)
	}
	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = dp[i+1]
		dp[i][s[i]-'a'] = i
	}

	// 遍历数组
	res := ""
	for _, t := range dictionary {
		i := 0 // s中的字符下标
		for _, c := range t {
			// 若c在s中没有出现过则跳出
			if dp[i][c-'a'] == len(s) {
				i = -1 // 当前t不是目标字符串，将i置为-1
				break
			}
			// 若出现过则将i移动到出现位置的下一个下标
			i = dp[i][c-'a'] + 1
		}

		// 若找到目标字符串且t的长度大于res或长度相等但t的字典序小于res，则将t替代res
		if i != -1 && (len(t) > len(res) || len(t) == len(res) && t < res) {
			res = t
		}
	}

	return res
}
