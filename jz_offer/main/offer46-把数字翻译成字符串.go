package main

import (
	"strconv"
)

// 剑指 Offer 46. 把数字翻译成字符串

// 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，
// ……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。
// 请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

func main() {

}

func translateNum(num int) int {
	numStr := strconv.Itoa(num)
	result := 1 // 总数
	pre1 := 1   // 前一个
	pre2 := 1   // 前两个

	str := ""

	for i := 1; i < len(numStr); i++ {
		str = numStr[i-1 : i+1]
		if str >= "10" && str <= "25" {
			result = pre1 + pre2
		} else {
			result = pre1
		}
		pre2 = pre1
		pre1 = result
	}
	return result
}

// 第二次做 -- 动态规划，类似于青蛙跳台阶，只是跳两级台阶时有条件了，当前数字和前面数字组合后的数字必须是在10~25之间
func translateNum1(num int) int {
	if num < 0 {
		return -1
	}

	// 先将数字转成字符串
	strNum := strconv.FormatInt(int64(num), 10)

	num1 := 1 // 前一项
	num2 := 1 // 前两项
	res := 1  // 第n项的总数

	for i := 1; i < len(strNum); i++ { // 从第二个字母开始计算
		temp := strNum[i-1 : i+1]
		if temp >= "10" && temp <= "25" {
			// 若和前一个数字组合后可以翻译为字母，则翻译当前数字可以和前面数字组合翻译，也可以单独翻译
			res = num1 + num2
		} else {
			// 若当前数字只能单独翻译
			res = num1
		}
		num2 = num1
		num1 = res
	}
	return res
}

func translateNum2(num int) int {
	if num < 0 {
		return -1
	}

	str := strconv.Itoa(num)
	pre1, pre2 := 1, 1
	res := 0
	for i := 1; i < len(str); i++ {
		tmp := str[i-1 : i+1]
		if tmp > "10" && tmp < "23" {
			res = pre1 + pre2
		} else {
			res = pre1
		}

		pre2 = pre1
		pre1 = res
	}
	return res
}
