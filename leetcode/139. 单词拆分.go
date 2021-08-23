package main

// 139. 单词拆分

func main() {

}

// 动态规划
func wordBreak(s string, wordDict []string) bool {
	if s == "" || wordDict == nil || len(wordDict) == 0 {
		return false
	}

	// 将单词放入map中
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	dp := make([]bool, len(s)+1) // dp[i]记录前i个字母能否拆成单次
	dp[0] = true

	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
