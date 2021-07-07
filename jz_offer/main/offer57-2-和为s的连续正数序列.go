package main

import (
	"fmt"
)

// 剑指 Offer 57 - II. 和为s的连续正数序列

// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
// 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

func main() {
	fmt.Println(findContinuousSequence1(9))
}

func findContinuousSequence(target int) [][]int {
	result := make([][]int, 0)
	small := 1 // 小数
	big := 2   // 大数

	res := make([]int, 0)
	sum := small + big
	for small < big {
		if sum > target {
			sum -= small
			small++
		} else if sum < target {
			big++
			sum += big
		} else {
			for s, b := small, big; s <= b; s++ {
				res = append(res, s)
			}
			result = append(result, res)
			res = make([]int, 0)

			big++
			sum += big
		}
	}
	return result
}

// 第二次做
// 使用双指针
func findContinuousSequence1(target int) [][]int {
	result := make([][]int, 0)
	l := 0
	r := 1

	sum := l + r // 当前的和
	for l < r {
		if sum > target {
			sum -= l
			l++
		} else if sum < target {
			r++
			sum += r
		} else {
			res := make([]int, r-l+1)
			left := l
			for i := 0; i < len(res); i++ {
				res[i] = left
				left++
			}
			result = append(result, res)

			r++
			sum += r
		}
	}
	return result
}
