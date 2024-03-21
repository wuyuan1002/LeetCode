package main

// 461. 汉明距离

// 两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
// 给你两个整数 x 和 y，计算并返回它们之间的汉明距离。

// hammingDistance .
// 先将两数字进行异或，异或的结果中二进制1的个数就是两数字中二进制位不同的位数
func hammingDistance(x int, y int) int {
	result := 0

	// 统计两数字异或后二进制中1的个数 -- 同 leetcode 191. 位1的个数
	for s := x ^ y; s != 0; s &= s - 1 {
		result++
	}

	return result
}
