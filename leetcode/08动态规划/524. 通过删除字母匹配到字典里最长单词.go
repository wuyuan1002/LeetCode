package main

// 524. 通过删除字母匹配到字典里最长单词

// 给你一个字符串 s 和一个字符串数组 dictionary ，找出并返回 dictionary 中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。
//
// 如果答案不止一个，返回长度最长且字母序最小的字符串。如果答案不存在，则返回空字符串。

// findLongestWord .
// 同 leetcode 1143. 最长公共子序列
//
// 1. 双指针遍历每一个元素和给定字符串，返回匹配的最长串 -- 会有大量的时间用于在s中找到下一个匹配字符
// 2. 动态规划对于s的每一个i位置，保存从该位置开始往后每一个字符第一次出现的位置
// 3. 对dictionary的每一个字符串t求 leetcode 1143. 最长公共子序列，若公共子序列长度正好等于t长度，说明t可以通过s删除一些字符得到，然后返回所有满足条件的t中最合适的那个即可
//
// dp[i][j]表示s中从位置i开始往后（即 s[i, len(s)) 范围内）字符j第一次出现的位置, 若 dp[i][j] == len(s) 则表示从位置i开始往后不存在字符j
func findLongestWord(s string, dictionary []string) string {

	// 构造dp数组，用来记录每个字符出现的位置 -- dp[i][j]表示s中从位置i开始往后字符j第一次出现的位置
	dp := make([][26]int, len(s)+1)

	// dp数组初始化 -- 初始化每个位置i之后每个字符j第一次出现的下标
	for i := range dp[len(s)] {
		dp[len(s)][i] = len(s)
	}
	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = dp[i+1]
		dp[i][s[i]-'a'] = i
	}

	// 开始dp -- 遍历每一个字符串，判断是否可以通过s删除字符得到，若可以则更新最优结果
	result := ""
	for _, t := range dictionary {
		// s中当前与t的相同字符下标
		i := 0

		// 遍历t中的每一个字符，判断是否在s的i之后出现过，若没出现过则说明t不能由s删除字符得到
		for _, c := range t {
			if dp[i][c-'a'] == len(s) {
				// 将i置为-1表示当前t不符合条件
				i = -1
				break
			}

			// 若t中当前字符c出现过则将i移动到s中出现位置的下一个下标
			i = dp[i][c-'a'] + 1
		}

		// 若找到目标字符串且t的长度大于result或长度相等但t的字典序小于result，则将t替代result
		if i != -1 && (len(t) > len(result) || len(t) == len(result) && t < result) {
			result = t
		}
	}

	return result
}
