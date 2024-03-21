package main

// 647. 回文子串

// 给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
// 回文字符串 是正着读和倒过来读一样的字符串。
// 子字符串 是字符串中的由连续字符组成的一个序列。
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

// countSubstrings .
// 同 leetcode 5、516
// 1. 双指针 -- 中心扩散法
// 2. 动态规划
//
// dp[i][j]表示s[i]与s[j]之间的子串是否为回文串（[i, j]为左闭右闭）
// dp[i][j] =
// 若s[i] != s[j]：dp[i][j] = false
// 若s[i] == s[j]：
//   1：下标i与j相同，同一个字符例如a，是回文子串 -- dp[i][j] = true
//   2：下标i与j相差为1，例如aa，是回文子串 -- dp[i][j] = true
//   3：下标i与j相差大于1的时候，例如cabac，此时s[i]与s[j]都指向c已经相同了，判断[i, j]是否为回文串，要依赖于[i+1, j-1]区间（即aba）是否为回文串（即dp[i+1][j-1]是否为true） -- dp[i][j] = dp[i + 1][j - 1]
//
// + j一定是大于等于i的，那么在填充dp[i][j]的时候其实是只填充右上半部分
// + 因为dp[i][j]依赖于dp[i+1][j-1], 因此dp数组应该是从左下往右上遍历 -- 字符串要倒着遍历
func countSubstrings(s string) int {

	// 总结果 -- 回文子串的总数量
	result := 0

	// 构造dp数组 -- dp[i][j]表示s[i]与s[j]之间的子串是否为回文串（[i, j]为左闭右闭）
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	// 初始化dp数组 -- 每个位置dp[i][j]都应该为false，表示每个字符
	// 因为bool默认为false，所以这里不用单独遍历dp进行初始化false的操作

	// 因为dp[i][j]依赖于dp[i+1][j-1], 因此dp数组应该是从左下往右上遍历 -- 字符串i要倒着遍历
	for i := len(s) - 1; i >= 0; i-- {
		// j遍历区间 [i, len(s)) 的所有字符，不断计算[i, j]区间的字符是否为回文串并累加result
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				// 若i和j指向的字符相等且满足以上的3中情况则说明[i, j]区间为回文串
				result++
				dp[i][j] = true
			}
		}
	}

	// 返回总结果
	return result
}
