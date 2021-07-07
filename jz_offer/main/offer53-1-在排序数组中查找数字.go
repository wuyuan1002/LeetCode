package main

import (
	"fmt"
)

// 剑指 Offer 53 - I. 在排序数组中查找数字 I

// 统计一个数字在排序数组中出现的次数。

func main() {
	// fmt.Println(search2([]int{1, 1}, 1))
	fmt.Println(search3([]int{1}, 1))
	// fmt.Println(search3([]int{1, 2, 3, 3, 3, 3, 3, 4, 6}, 3))
}

// 方法1：因为数组是排好序的，所以可以使用二分法查找第一个k和最后一个k，最后下标的差值就是出现的次数
func search1(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	num := 0
	first := getFirstK(&nums, target, 0, len(nums)-1)
	last := getLastK(&nums, target, 0, len(nums)-1)

	if first != -1 && last != -1 {
		num = last - first + 1
	}
	return num
}

// 获取第一个k的下标
func getFirstK(nums *[]int, k, left, right int) int {
	if left > right {
		return -1
	}

	middle := (left + right) / 2 // 中间元素下标
	if (*nums)[middle] > k {     // 若中间元素比k大，说明k在前半部分
		right = middle - 1
	} else if (*nums)[middle] < k { // 若中间元素比k小，说明k在后半部分
		left = middle + 1
	} else { // 若中间元素等于k，查看中间元素的前一个数字是不是也等于k，若不等于，则中间元素就是第一个k，若等于，说明中间元素不是第一个k，第一个k还在前半部分
		if (middle > 0 && (*nums)[middle-1] != k) || middle == 0 {
			return middle
		} else {
			right = middle - 1
		}
	}

	return getFirstK(nums, k, left, right)
}

// 获取最后一个k的下标
func getLastK(nums *[]int, k, left, right int) int {
	if left > right {
		return -1
	}

	middle := (left + right) / 2 // 中间元素下标
	if (*nums)[middle] > k {     // 若中间元素比k大，说明k在前半部分
		right = middle - 1
	} else if (*nums)[middle] < k { // 若中间元素比k小，说明k在后半部分
		left = middle + 1
	} else { // 若中间元素等于k，查看中间元素的后一个数字是不是也等于k，若不等于，则中间元素就是后一个k，若等于，说明中间元素不是最后一个k，最后一个k还在后半部分
		if (middle < len(*nums)-1 && (*nums)[middle+1] != k) || middle == len(*nums)-1 {
			return middle
		} else {
			left = middle + 1
		}
	}

	return getLastK(nums, k, left, right)
}

// 方法2：把数字存入map，键是数字、值是出现的次数
func search2(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	numMap := make(map[int]int)
	for _, num := range nums {
		if n, ok := numMap[num]; ok { // 若已存在，把数量加1
			numMap[num] = n + 1
		} else {
			numMap[num] = 1
		}
	}

	return numMap[target]
}

// 第二次做 -- 二分法
func search3(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	num := 0
	first := getFirstK1(nums, target, 0, len(nums)-1)
	last := getLastK1(nums, target, 0, len(nums)-1)

	if first != -1 && last != -1 {
		num = last - first + 1
	}
	return num
}

// 获取第一个k的下标
func getFirstK1(nums []int, k, left, right int) int {

	for left <= right {
		middle := (left + right) / 2 // 中间元素下标
		if nums[middle] > k {
			right = middle - 1
		} else if nums[middle] < k {
			left = middle + 1
		} else {
			if middle > 0 && nums[middle-1] == k {
				right = middle - 1
			} else {
				return middle
			}
		}
	}

	return -1
}

// 获取最后一个k的下标
func getLastK1(nums []int, k, left, right int) int {

	for left <= right {
		middle := (left + right) / 2 // 中间元素下标
		if nums[middle] > k {
			right = middle - 1
		} else if nums[middle] < k {
			left = middle + 1
		} else {
			if middle < len(nums)-1 && nums[middle+1] == k {
				left = middle + 1
			} else {
				return middle
			}
		}
	}

	return -1
}
