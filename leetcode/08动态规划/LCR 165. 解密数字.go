package main

import "strconv"

// LCR 165. 解密数字

// 现有一串神秘的密文 ciphertext，经调查，密文的特点和规则如下：
//
// 密文由非负整数组成
// 数字 0-25 分别对应字母 a-z
// 请根据上述规则将密文 ciphertext 解密为字母，并返回共有多少种解密结果。

// crackNumber .
// 同 leetcode 91. 解码方法、509. 斐波那契数、70. 爬楼梯、198. 打家劫舍
//
// dp[i]表示以s[i]结尾的子串的解码方法个数
// dp[i] = dp[i-1] -- s[i]不能与前一个字符组合进行解码构成字母
// dp[i] = dp[i-1] + dp[i-2] -- s[i]可以与前一个字符组合进行解码构成字母
func crackNumber(ciphertext int) int {
	s := strconv.Itoa(ciphertext)
	if len(s) == 1 {
		return 1
	}

	// 构造dp数组 -- dp[i]表示以s[i]结尾的子串的解码方法个数
	dp := make([]int, len(s))

	// 初始化dp数组 -- dp[0]表示第一个字符只有一种解密方式，dp[1]若第二个字符可以和前一个字符构成字母则有两种解密方式，若构不成则有一种解密方式
	dp[0] = 1
	if (s[0]-'0')*10+(s[1]-'0') < 26 {
		dp[1] = 2
	} else {
		dp[1] = 1
	}

	// 开始dp -- 求s的第2个到最后一个字符结尾的解码个数
	for i := 2; i < len(s); i++ {
		if s[i-1] != '0' && ((s[i-1]-'0')*10+(s[i]-'0') < 26) {
			// 当前位置字符可以与前一个字符组合进行解码构成字母
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			// 当前位置字符无法与前一个字符组合进行解码构成字母
			dp[i] = dp[i-1]
		}
	}

	return dp[len(s)-1]
}
