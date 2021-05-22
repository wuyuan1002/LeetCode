package main

import "fmt"

// 剑指 Offer 14- I. 剪绳子

// 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），
// 每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
// 例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
func main() {
	fmt.Printf("%d", cuttingRope1(10))
}

// 动态规划算法的核心就是记住已经解决过的子问题的解 -- 见offer42

// 自底向上的动态规划 -- 先算小的，后算大的
func cuttingRope(n int) int {
	if n < 2 { // 只有一种剪的方式，直接返回
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	products := make([]int, n+1) // 存放已经计算好的小长度的最优解
	products[0] = 0              // 最优子结构 -- 下标为绳子长度，到这么长就别切了，因为切了不如不切
	products[1] = 1              // 最优子结构 -- 下标为绳子长度，到这么长就别切了，因为切了不如不切
	products[2] = 2              // 最优子结构 -- 下标为绳子长度，到这么长就别切了，因为切了不如不切
	products[3] = 3              // 最优子结构 -- 下标为绳子长度，到这么长就别切了，因为切了不如不切

	max := 0 // 当前问题的最优解 -- 当前长度切的话乘积的最大值
	for i := 4; i <= n; i++ {
		max = 0
		for j := 1; j <= i/2; j++ {
			product := products[j] * products[i-j] // 从已经算好的子问题中求解当前问题的解
			if max < product {
				max = product
			}
			products[i] = max // 将当前问题的最优解存起来，让求解大问题时使用
		}
	}
	max = products[n]
	return max
}

// 第二次做
// 动态规划
func cuttingRope1(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	products := make([]int, n+1) // 保存已有的解

	// 最优子结构
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3

	max := 0     // 某个长度乘积的最大值
	product := 0 // 某次截取的乘积
	for i := 4; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			product = products[j] * products[i-j]
			if product > max {
				max = product
			}
		}
		products[i] = max
		max = 0
	}
	return products[n]
}
