package main

// 青蛙跳台阶问题

//  一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n级的台阶总共有多少种跳法。
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
func main() {

}

func numWays(n int) int {
	if n < 2 {
		return 1
	}

	a, b := 1, 1
	for i := 2; i <= n; i++ {
		temp := (a + b) % constant
		a = b
		b = temp
	}
	return b
}

// 第二次做 -- 动态规划，后一项等于前两项的和
func numWays1(n int) int {
	if n < 2 {
		return 1
	}

	num1 := 1 // 一级前的跳法数 -- 第n项的前一项
	num2 := 1 // 两级前的跳法数 -- 第n项的前两项
	num := 1  // 总共有多少种跳法

	for i := 2; i <= n; i++ { // 从第二级台阶开始算
		num = (num1 + num2) % constant
		num2 = num1
		num1 = num
	}
	return num
}
