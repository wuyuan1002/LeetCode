package main

import "math"

// 453. 最小操作次数使数组元素相等

// 给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。

// func main() {

// }

// 每次给n-1个数减1相当于每次给一个数加1
// 因此只需要计算出让所有元素与数组中最小元素相等时的操作次数即可
func minMoves(nums []int) int {
	// 求出数组中的最小值
	min := math.MaxInt32
	for _, n := range nums {
		if n < min {
			min = n
		}
	}

	// 	计算每个元素与最小值的差
	res := 0
	if min != math.MaxInt32 {
		for _, n := range nums {
			res += n - min
		}
	}

	return res
}
