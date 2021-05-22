package main

import (
	"fmt"
	"strconv"
)

// 剑指 Offer 17. 打印从1到最大的n位数

// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。
// 比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
func main() {
	fmt.Printf("%v", printNumbers1(3))
}

// 不推荐 -- 大数问题容易越界
func printNumbers(n int) []int {
	number := 1
	count := 1
	for i := 1; i <= n; i++ {
		count *= 10
	}
	results := make([]int, count-1)
	for ; number < count; number++ {
		results[number-1] = number
	}
	return results
}

// 第二次做 -- 使用byte数组存储数字，解决大数问题
func printNumbers1(n int) []int {
	count := 1
	for i := 1; i <= n; i++ {
		count *= 10
	}
	results := make([]int, count-1) // 存放所有结果

	number := make([]byte, n) // 存放当前数字
	printN(&number, &results, n, 0)
	return results
}

// length表示数字的总长度，index表示现在是数字的第几位
func printN(number *[]byte, results *[]int, length, index int) {

	// 如果已经到最后一位，将byte数组转成数字存入切片中
	if index == length {
		str := string(*number)
		if n, err := strconv.Atoi(str); err == nil && n != 0 {
			(*results)[n-1] = n
		}
		return
	}

	for i := 0; i < 10; i++ { // 把当前位从0~9的每个数字都放一遍
		(*number)[index] = byte(i + '0')
		printN(number, results, length, index+1)
	}
}
