package main

import (
	"fmt"
)

// 剑指 Offer 53 - II. 0～n-1中缺失的数字

// 一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
// 在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

func main() {
	fmt.Println(missingNumber([]int{0}))
}

// 数组递增排序，且数字范围都在n-1范围内，说明缺失的数字正好是第一个与下标不相等的数字
// 使用二分法
func missingNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == middle { // 中间元素相等，说明缺失的数字还在右面
			left = middle + 1
		} else { // 若不相等
			if middle == 0 || nums[middle-1] == middle-1 { // 若前一个数字也不想等或中间元素已经是第一个数字了，说明这个数字就是缺失的那个
				return middle
			}
			right = middle - 1 // 若前一个数字也不相等，说明缺失的数字还在左面
		}
	}
	if left == len(nums) {
		return len(nums)
	}
	return -1
}

// 第二次做
// 使用二分法 -- 查找第一个与下标不相等的数字
func missingNumber1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	l := 0
	r := len(nums) - 1
	for l <= r {
		middle := (l + r) / 2
		if nums[middle] == middle {
			l = middle + 1
		} else {
			if middle > 0 && nums[middle-1] != middle-1 {
				r = middle - 1
			} else {
				return middle
			}
		}
	}
	return len(nums)
}
