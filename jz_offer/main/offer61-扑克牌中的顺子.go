package main

import (
	"fmt"
)

// 剑指 Offer 61. 扑克牌中的顺子

// 从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，
// A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

func main() {
	fmt.Println(isStraight([]int{1, 0, 2, 0, 5}))
}

func isStraight(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}

	// 排序
	quickSort(nums, 0, len(nums)-1)

	lastNum := 0 // 记录上一个数字
	zeroNum := 0 // 0的个数
	for _, v := range nums {
		if v == 0 { // 统计0的个数
			zeroNum++
		} else if lastNum != 0 && v == lastNum { // 有对子肯定不是顺子 -- 有相等的数字 -- 0123345
			return false
		} else if lastNum != 0 && v != lastNum+1 { // 相邻两个数字不连续 -- 用0补齐，若0的个数不足，则不能构成顺子
			gap := v - lastNum - 1
			if gap > zeroNum {
				return false
			}
			zeroNum -= gap
		}

		lastNum = v // 记录当前数字，循环到下一个数字时看是不是相等 -- 有没有对子
	}

	return true
}

func quickSort(arr []int, left int, right int) {

	if left >= right || arr == nil || len(arr) == 0 {
		return
	}

	l, r := left, right
	temp := arr[left]
	for l < r {
		for arr[r] >= temp && l < r {
			r--
		}
		for arr[l] <= temp && l < r {
			l++
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	arr[left], arr[l] = arr[l], temp

	quickSort(arr, left, l-1)
	quickSort(arr, r+1, right)
}
