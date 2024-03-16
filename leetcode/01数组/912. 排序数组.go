package main

// 912. 排序数组

// 给你一个整数数组 nums，请你将该数组升序排列。

// sortArray .
// https://mp.weixin.qq.com/s/5HqfRGqPyAhFt0krPgMHOQ
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// quickSort 递归快排
func quickSort(nums []int, left, right int) {
	if nums == nil || len(nums) == 0 || left >= right {
		return
	}

	l, r := left, right
	temp := nums[left]
	for l < r {
		for l < r && nums[r] >= temp {
			r--
		}
		for l < r && nums[l] <= temp {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[left], nums[l] = nums[l], nums[left]

	quickSort(nums, left, l-1)
	quickSort(nums, r+1, right)
}
