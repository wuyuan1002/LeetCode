package main

import (
	"fmt"
)

// 剑指 Offer 60. n个骰子的点数

// 把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
//
// 你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。

func main() {
	fmt.Println(dicesProbability(2))
}

// 动态规划，tmp[x+y] = pre[x]*num[y]
// 对掷骰子，num永远是1/6，pre初始化为第一次掷骰子的结果
func dicesProbability(n int) []float64 {

	// 最优子结构 -- 一个骰子的情况
	res := []float64{1.0 / 6, 1.0 / 6, 1.0 / 6, 1.0 / 6, 1.0 / 6, 1.0 / 6}
	// 第n个骰子的情况
	var tmp []float64

	// 从第2个骰子开始计算
	for i := 2; i <= n; i++ {
		tmp = make([]float64, 5*i+1) // 第n个骰子投完后各个数量的概率
		for k := 0; k < 6; k++ {     // 第n个骰子分别投出 1~6 点数时
			for j := 0; j < len(res); j++ { // 计算每一个点数对应的新的概率
				tmp[j+k] += res[j] / 6
			}
		}
		res = tmp
	}
	return res
}
