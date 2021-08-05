package main

import "fmt"

// 713. 乘积小于K的子数组

// 给定一个正整数数组 nums和整数 k 。
// 请找出该数组内乘积小于 k 的连续的子数组的个数。

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{1, 50, 3, 6}, 100))
}

// 双指针
func numSubarrayProductLessThanK(nums []int, k int) int {
	if nums == nil || len(nums) == 0 || k <= 1 {
		return 0
	}

	// 左指针，当前的乘积，结果
	left, curAccu, result := 0, 1, 0
	for right, v := range nums {
		curAccu *= v
		for curAccu >= k {
			curAccu /= nums[left]
			left++
		}
		result += right - left + 1
	}
	return result
}
