package main

import "fmt"

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面

// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，
// 所有偶数位于数组的后半部分。

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%v", exchange(nums))
}

func exchange(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	i := 0
	j := len(nums) - 1

	for i < j {
		for ; i < j && nums[j]&0x1 == 0; j-- {
		}
		for ; i < j && nums[i]&0x1 == 1; i++ {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return nums
}

// 第二次做 -- 双指针法
func exchange1(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nums
	}

	i := 0
	j := len(nums) - 1
	for i < j {
		for ; i < j && isEven(nums[i]); i++ {
		}
		for ; j > i && !isEven(nums[j]); j-- {
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}

func isEven(num int) bool {
	return num&0x1 == 1 // 判断是不是奇数
}
