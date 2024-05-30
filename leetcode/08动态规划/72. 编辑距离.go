package main

// 72. 编辑距离

// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
//
// 你可以对一个单词进行如下三种操作：
//
// 插入一个字符
// 删除一个字符
// 替换一个字符

// minDistance .
// 同 leetcode 583. 两个字符串的删除操作
//
// dp[i][j]表示以下标i-1结尾的word1变换到下标j-1结尾的word2最少需要操作的次数
//
// dp[i-1][j-1]表示替换:
// 	 word1的前i-1个字符已经变换到word2的前j-1个字符的次数，说明word1的前i-1个和word2的前j-1个字符已经完成操作；
// 	 那么对于word1的第i个怎么变成word2的第j个呢？这两个字符都存在，那么只能是替换了 -- dp[i][j] = dp[i-1][j-1]+1
// dp[i][j-1]表示插入:
// 	 word1的前i个字符已经变换到word2的前j-1个字符的次数，当前word1的第i步字符已经用了，
// 	 但是word2还差一个字符（因为当前只是处理了word2的前j-1个字符），那么插入一个字符就好了 -- dp[i][j] = dp[i][j-1]+1
// dp[i-1][j]表示删除:
// 	 word1的前i-1个字符已经变换到word2的前j个字符的次数，当前word1仅用了前i-1个字符就完成了到word2的前j个字符的操作，
// 	 所以word1的第i个字符其实没啥用了，那么删除操作就好了 -- dp[i][j] = dp[i-1][j]+1
func minDistance(word1 string, word2 string) int {

	// 构造dp数组 -- dp[i][j]表示以下标i-1结尾的word1变换到下标j-1结尾的word2最少需要操作的次数
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	// 初始化dp数组 -- 第一行表示空字符串word1变到word2需要几步，第一列表示word1变到空字符串word2需要几步
	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(word2)+1; j++ {
		dp[0][j] = j
	}

	// 开始dp -- 遍历每个位置，计算该位置从word1到word2需要几步
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			left_top := dp[i-1][j-1]
			// 若当前位置的前两字符不相等则执行替换操作
			if word1[i-1] != word2[j-1] {
				left_top++
			}

			// 计算当前位置的最短编辑距离 -- min(执行替换操作, min(执行插入操作, 执行删除操作))
			dp[i][j] = min(left_top, min(dp[i][j-1]+1, dp[i-1][j]+1))
		}
	}

	return dp[len(word1)][len(word2)]
}
