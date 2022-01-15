package main

// 剑指 Offer 49. 丑数

// 我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

// func main() {

// }

// 动态规划 -- 当前丑数是由我前面的丑数x2、x3、x5得到的
func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	x2, x3, x5 := 0, 0, 0 // 下一个需要 x2、x3、x5 的数字下标
	products := make([]int, n)
	products[0] = 1
	for i := 1; i < n; i++ {
		n2, n3, n5 := products[x2]*2, products[x3]*3, products[x5]*5
		products[i] = min(min(n2, n3), n5)
		if products[i] == n2 {
			x2++ // 下一个应该x2的数字
		}
		if products[i] == n3 {
			x3++ // 下一个应该x3的数字
		}
		if products[i] == n5 {
			x5++ // 下一个应该x5的数字
		}
	}
	return products[n-1]
}

// 第二次做
// 动态规划 -- 当前丑数是由我前面的丑数x2、x3、x5得到的
func nthUglyNumber1(n int) int {
	if n <= 0 {
		return 0
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	uglyNumbers := make([]int, n)
	uglyNumbers[0] = 1 // 最优子结构 -- 第一个丑数是1

	x2, x3, x5 := 0, 0, 0 // 下一个需要 x2、x3、x5 的数字下标 -- 保证每个丑数只会被乘3次
	for i := 1; i < n; i++ {
		n2, n3, n5 := uglyNumbers[x2]*2, uglyNumbers[x3]*3, uglyNumbers[x5]*5
		uglyNumbers[i] = min(min(n2, n3), n5) // 获取乘完之后最小的那个作为下一个丑数
		if uglyNumbers[i] == n2 {
			x2++
		}
		if uglyNumbers[i] == n3 {
			x3++
		}
		if uglyNumbers[i] == n5 {
			x5++
		}
	}
	return uglyNumbers[n-1]
}
