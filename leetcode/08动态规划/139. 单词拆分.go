package main

// 139. 单词拆分

// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
//
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

// wordBreak .
// 动态规划 -- 完全背包
// 单词就是物品，字符串s就是背包，单词能否组成字符串s，就是问物品能不能把背包装满 -- 也就是字符串s能否被选择列表中的单词拼接构成
// 拆分时可以重复使用字典中的单词，说明是一个完全背包
// dp[i]表示字符串的前i部分能否被选择列表中的单词拼接构成
func wordBreak(s string, wordDict []string) bool {
	// 将单词加入map中 -- 提供O(1)的查询时间复杂度
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	// 构造dp数组
	dp := make([]bool, len(s)+1)
	// 初始化dp数组 -- 无意义
	dp[0] = true

	// 开始dp -- 计算字符串前i部分能否被选择列表中的单词拼接构成
	// 本题求的是排列数，所以应该先遍历背包再遍历选择列表 -- 因为物品的顺序影响结果，所以是排列问题 leetcode 377
	for i := 1; i <= len(s); i++ { // 遍历背包容量 -- 遍历字符串s
		for j := 0; j < i; j++ { // 遍历选择列表 -- 若s的[0, j)和[j, i)都可以被选择列表中的单词构成，那么[0, i)部分也就可以被选择列表中的单词构成
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	// 返回字符串s能否可以被选择列表中的单词拼接构成
	return dp[len(s)]
}
