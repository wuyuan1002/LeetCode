package main

import (
	"fmt"
)

// 215. 数组中的第K个最大元素

// 在未排序的数组中找到第 k 个最大的元素。
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

// 同 offer 40
func findKthLargest(nums []int, k int) int {
	if nums == nil || len(nums) == 0 || k <= 0 || k >= len(nums) {
		return 0
	}

	left, right := 0, len(nums)-1
	index := partition(nums, left, right)
	for index != k-1 {
		if index > k-1 {
			index = partition(nums, left, index-1)
		} else {
			index = partition(nums, index+1, right)
		}
	}

	return nums[index]
}

func partition(nums []int, left, right int) int {

	l, r := left, right
	tmp := nums[left]

	for l < r {
		for l < r && nums[r] <= tmp {
			r--
		}
		for l < r && nums[l] >= tmp {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

	nums[left], nums[l] = nums[l], tmp
	return l
}
