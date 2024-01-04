package main

// 338. 比特位计数

// 给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，
// 返回一个长度为 n + 1 的数组 ans 作为答案。

// countBits .
// 把一个整数减1再和原来的整数作与运算，会把该整数最右面的1变成0 -- 一个整数中有多少个1就可以做多少次这样的操作
func countBits(n int) []int {
	// 总结果 -- 记录每个数字中1的个数
	result := make([]int, n+1)

	// 遍历每一个数字，计算该数字中1的个数并保存到总结果中
	for i := 1; i <= n; i++ {
		num := i
		// 不断将当前数字与它减一的数字做与运算并累加1的个数
		for num != 0 {
			result[i]++
			num &= num - 1
		}
	}

	return result
}
