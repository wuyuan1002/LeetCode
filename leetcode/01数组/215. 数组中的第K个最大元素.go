package main

// 215. 数组中的第K个最大元素

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

// findKthLargest .
// leetcode 347. 前 K 个高频元素
//
// 使用快排partition -- 在快排的过程中找到目标值后停止排序直接返回即可
func findKthLargest(nums []int, k int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	// 不断归位一个基准数，并判断该归位的数字是否为第k个，若不是则继续归位
	index := partition(nums, 0, len(nums)-1)
	for index != k-1 {
		if index > k-1 {
			// 若当前归位的基准数下标比k大，说明第k大的元素一定位于当前元素的左面，所以继续归位nums[0, index-1]即可
			index = partition(nums, 0, index-1)
		} else {
			// 若当前归位的基准数下标比k小，说明第k大的元素一定位于当前元素的右面，所以继续归位nums[index+1, len(nums)-1]即可
			index = partition(nums, index+1, len(nums)-1)
		}
	}

	return nums[index]
}

// partition 降序快排的一次partition，归位一个基准数，并返回基准数归位后的下标
func partition(nums []int, left, right int) int {
	l, r := left, right // 左右指针
	temp := nums[left]  // 选择最左面数字作为基准数

	for l < r {
		for l < r && nums[r] <= temp {
			r--
		}
		for l < r && nums[l] >= temp {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[left], nums[l] = nums[l], nums[left]

	return l
}
