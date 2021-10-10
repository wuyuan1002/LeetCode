package main

// 441. 排列硬币

// 你总共有n枚硬币，并计划将它们按阶梯状排列。对于一个由 k 行组成的阶梯，
// 其第 i 行必须正好有 i 枚硬币。阶梯的最后一行 可能 是不完整的。
//
// 给你一个数字n ，计算并返回可形成 完整阶梯行 的总行数。

func main() {

}

// 等差数列求和 -- 首项加末项乘以项数除以2
func arrangeCoins(n int) int {
	if n <= 0 {
		return 0
	} else if n < 3 {
		return 1
	}

	// i: 项数 -- 第几行
	// 这里的i可以使用二分法在1到n中查找
	for i := 2; ; i++ {
		if sum := (1 + i) * i / 2; sum == n {
			return i
		} else if sum > n {
			return i - 1
		}
	}
}
