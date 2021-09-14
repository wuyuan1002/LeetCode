package main

// 524. 通过删除字母匹配到字典里最长单词

// 给你一个字符串 s 和一个字符串数组 dictionary 作为字典，找出并返回字典中最长的字符串，
// 该字符串可以通过删除 s 中的某些字符得到。
//
// 如果答案不止一个，返回长度最长且字典序最小的字符串。如果答案不存在，则返回空字符串。

func main() {
	findLongestWord1("abcde", []string{"aplpllre"})
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
		for _, c := range t {
			if dp[c-'a'] == len(s) {

			}
		}
	}
}

func findLongestWord1(s string, dictionary []string) (ans string) {
	m := len(s)
	f := make([][26]int, m+1)
	for i := range f[m] {
		f[m][i] = m
	}
	for i := m - 1; i >= 0; i-- {
		f[i] = f[i+1]
		f[i][s[i]-'a'] = i
	}

outer:
	for _, t := range dictionary {
		j := 0
		for _, ch := range t {
			if f[j][ch-'a'] == m {
				continue outer
			}
			j = f[j][ch-'a'] + 1
		}
		if len(t) > len(ans) || len(t) == len(ans) && t < ans {
			ans = t
		}
	}
	return
}
