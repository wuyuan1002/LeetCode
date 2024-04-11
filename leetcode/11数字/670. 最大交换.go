package main

import "strconv"

// 670. 最大交换

// 给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

// maximumSwap .
// 将右边越大的数字与左边较小的数字进行交换，这样产生的整数才能保证越大
func maximumSwap(num int) int {
	// 将数字每一位转换成字符
	s := []byte(strconv.Itoa(num))

	// 当前遍历到的最大值下标、右侧最大的数字下标、左侧第一个小于右侧最大数字的数字下标
	maxIdx, rightMaxIdx, leftMinIdx := len(s)-1, -1, -1

	// 倒序遍历每一位数字，找到最优的被交换的两位下标
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] > s[maxIdx] {
			maxIdx = i
		} else if s[i] < s[maxIdx] {
			leftMinIdx, rightMaxIdx = i, maxIdx
		}
	}

	// 若左侧没有小于右侧最大数字的数字，说明当前数字已经是最大
	if leftMinIdx < 0 {
		return num
	}

	// 交换两下标对应数字并返回总结果
	s[leftMinIdx], s[rightMaxIdx] = s[rightMaxIdx], s[leftMinIdx]
	v, _ := strconv.Atoi(string(s))

	return v
}
