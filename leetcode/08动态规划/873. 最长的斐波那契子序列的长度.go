package main

// 873. 最长的斐波那契子序列的长度

// 如果序列 X_1, X_2, ..., X_n 满足下列条件，就说它是 斐波那契式 的：
//
// n >= 3
// 对于所有 i + 2 <= n，都有 X_i + X_{i+1} = X_{i+2}
// 给定一个严格递增的正整数数组形成序列 arr ，找到 arr 中最长的斐波那契式的子序列的长度。如果一个不存在，返回  0 。
//
// （回想一下，子序列是从原序列 arr 中派生出来的，它从 arr 中删掉任意数量的元素（也可以不删），
// 而不改变其余元素的顺序。例如， [3, 5, 8] 是 [3, 4, 5, 6, 7, 8] 的一个子序列）

// lenLongestFibSubseq .
// 根据题意可知，斐波那契子序列的长度至少为3，因此对于下标满足条件 k < i < j 的三个下标：
//
// dp[i][j]表示以arr[i]为倒数第二位，arr[j]为最后一位的最长斐波那契子序列长度
// dp[i][j] = max(dp[k][i]+1, 3) -- 当数组中存在下标k可以与i和j构成斐波那契数列时，即存在下标k满足 arr[k] + arr[i] == arr[j]
// dp[i][j] = 0 -- 当数组中不存在下标k可以与i和j构成斐波那契数列时
func lenLongestFibSubseq(arr []int) int {
	// 总结果 -- 最长的斐波那契子序列长度
	result := 0

	// 使用一个map索引数组数字与下标的对应关系，用来快速查找数组中是否存在等于某个值的下标k
	hash := make(map[int]int, len(arr))
	for i, x := range arr {
		hash[x] = i
	}

	// 构造dp数组 -- dp[i][j]表示以arr[i]为倒数第二位，arr[j]为最后一位的最长斐波那契子序列长度
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, len(arr))
	}

	// 初始化dp数组 -- 初始时以每个数字结尾的斐波那契数列长度为0

	// 开始dp -- 从小到大遍历下标i，从大到小遍历下标j，查找满足条件的下标k，更新dp数组的过程中更新总结果最大值
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j >= 0; j-- {
			// 若数组中存在下标k满足k < i < j，且 arr[k] + arr[i] == arr[j]，说明以arr[i]和arr[j]结尾的斐波那契数列左侧可以增添新的元素arr[k]了
			if k, ok := hash[arr[j]-arr[i]]; ok && k < i {
				// 更新dp数组，既然arr[k]可以作为以arr[i]和arr[j]结尾的斐波那契数列的元素，
				// 那么以arr[i]和arr[j]结尾的斐波那契数列长度就等于max(以arr[k]和arr[i]结尾的斐波那契数列长度 + 1, 3)
				dp[i][j] = max(dp[k][i]+1, 3)

				// 更新最长的数列长度
				result = max(result, dp[i][j])
			}
		}
	}

	return result
}
