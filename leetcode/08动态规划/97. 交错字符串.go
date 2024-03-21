package main

// 97. 交错字符串

// 给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。
//
// 两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
//
// s = s1 + s2 + ... + sn
// t = t1 + t2 + ... + tm
// |n - m| <= 1
// 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
// 注意：a + b 意味着字符串 a 和 b 连接。

// isInterleave .
// dp[i][j]表示s1的前i个字符s1[0, i)和s2的前j个字符s2[0, j)能否构成s3的前i+j个字符s3[0, i+j-1)
// 第一行：dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
// 第一列：dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
// 其它：dp[i][j] = (dp[i][j-1] && s2[j-1] == s3[i+j-1]) || (dp[i-1][j] && s1[i-1] == s3[i+j-1])
func isInterleave(s1 string, s2 string, s3 string) bool {
	// 若长度不相等则直接返回false
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	// 构造dp数组 -- dp[i][j]表示s1的前i个字符s1[0, i]和s2的前j个字符s2[0, j]能否构成s3的前i+j个字符s3[0, i+j]
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	// 初始化dp数组 -- dp[0][0]表示两个空串能否够构成空串，一定是true，其它位置为默认false
	dp[0][0] = true

	// 开始dp -- 求s1和s2中每一个下标的字符能否构成s3
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true // dp[0][0]表示两个空串能否够构成空串，一定是true
			} else if i == 0 { // 第一行
				dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
			} else if j == 0 { // 第一列
				dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
			} else {
				dp[i][j] = (dp[i][j-1] && s2[j-1] == s3[i+j-1]) || (dp[i-1][j] && s1[i-1] == s3[i+j-1])
			}
		}
	}

	return dp[len(s1)][len(s2)]
}
