package main

// 139. 单词拆分

// 给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，
// 判定s是否可以被空格拆分为一个或多个在字典中出现的单词。
//
// 说明：
//
// 拆分时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。

func main() {

}

// 动态规划
func wordBreak(s string, wordDict []string) bool {
	if s == "" || wordDict == nil || len(wordDict) == 0 {
		return false
	}

	// 将单词加入map中
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	// 记录前i个能否拆成单词
	products := make([]bool, len(s)+1)
	products[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if products[j] && wordMap[s[j:i]] {
				products[i] = true
				break
			}
		}
	}

	return products[len(s)]
}
