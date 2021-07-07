package main

// 70. 爬楼梯

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 注意：给定 n 是一个正整数。

func main() {

}

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}

	pre1, pre2 := 1, 1
	for i := 2; i <= n; i++ {
		num := pre1 + pre2
		pre2 = pre1
		pre1 = num
	}

	return pre1
}
