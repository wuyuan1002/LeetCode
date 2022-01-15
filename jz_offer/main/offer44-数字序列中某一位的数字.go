package main

import (
	"math"
)

// 剑指 Offer 44. 数字序列中某一位的数字

// 数字以0123456789101112131415…的格式序列化到一个字符序列中。
// 在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。
//
// 请写一个函数，求任意第n位对应的数字。

// func main() {
// 	// fmt.Println(findNthDigit(10))

// 	fmt.Println(getNumsCount(2))
// }

func findNthDigit(n int) int {
	if n < 0 {
		return -1
	}
	if n < 10 {
		return n
	}

	digits := 1 // 记录当前数字位数
	count := 0  // 跳过的总位数
	countNow := 0
	for {
		countNow = getNumsCount(digits)
		if n >= count+countNow {
			count += countNow
			digits++
			continue
		}
		break
	}

	indexAll := n - count         // 在所有digest位数字的总下标
	indexNum := indexAll % digits // 在确定的某一个digist数字内的下标

	startNum := math.Pow(10, float64(digits-1)) // digest位数字的开始数字

	num := int(startNum) + (indexAll / digits) // 所在的digits数字
	for i := 1; i < digits-indexNum; i++ {
		num /= 10
	}
	return num % 10
}

// 获取n位数字总共有几位
func getNumsCount(n int) int {
	if n == 1 {
		return 10
	}

	num := n // 记录数字位数
	count := 9
	for ; n > 1; n-- {
		count *= 10
	}
	return count * num
}
