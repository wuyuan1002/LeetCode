package main

import (
	"fmt"
)

// 剑指 Offer 56 - II. 数组中数字出现的次数 II

// 在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

func main() {
	// 方法1：算出每一位二进制位中1出现的次数，每一位次数与3取余
	// 方法2：使用map记录次数

	fmt.Println(singleNumber1([]int{13, 13, 13, 3}))
}

// 方法1：算出每一位二进制位中1出现的次数，每一位次数与3取余
func singleNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	bitNum := make([]int, 32) // 记录二进制位的每一位上1出现的次数 -- int为32位

	for _, n := range nums {
		mask := 1
		for i := 31; i >= 0; i-- {
			if n&mask != 0 { // 若第i位为1
				bitNum[i]++
			}
			mask <<= 1
		}
	}

	result := 0
	for i := 0; i < 32; i++ {
		result <<= 1
		result += bitNum[i] % 3
	}

	return result
}

// 第二次做
func singleNumber1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	bitNum := make([]int, 32)
	flag := 1
	for _, num := range nums { // 遍历每一个数字
		for i := 31; i > 0; i-- { // 遍历当前数字的每一位
			if num&flag != 0 {
				bitNum[i]++
			}
			flag <<= 1
		}
		flag = 1
	}

	result := 0
	for i := 0; i < 32; i++ {
		result <<= 1
		result += bitNum[i] % 3
	}
	return result
}
