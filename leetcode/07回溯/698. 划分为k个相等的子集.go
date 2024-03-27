package main

// 698. 划分为k个相等的子集

// 给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。

// canPartitionKSubsets .
// 1. 回溯法 -- 同 leetcode 40. 组合总和
// 2. 动态规划
func canPartitionKSubsets(nums []int, k int) bool {
	// 计算数组总和
	sum := 0
	for _, n := range nums {
		sum += n
	}
	// 若数组总和不是k的倍数，说明无法被平均分成k份，或数组中最大的数字已经比目标总和大了，那么直接返回false
	if sum%k != 0 || nums[len(nums)-1] > sum/k {
		return false
	}

	targetSum := sum / k            // 目标总和
	used := make([]bool, len(nums)) // 用来标记对应位置的数字是否已被使用

	// 将数组从小到大进行排序
	quickSort698(nums, 0, len(nums)-1)

	return dfsCanPartitionKSubsets(nums, len(nums)-1, targetSum, 0, k, used)
}

// dfsCanPartitionKSubsets .
// nums: 选择列表
// start: 开始下标 -- 从后往前遍历选择数字，因为需要先将大数进行匹配，防止前面的多个小数相加构成目标和了，导致剩下后面的大数没有小数与其相加构成目标和的情况，所以要先将大数进行匹配
// targetSum: 目标和
// currentSum: 当前和
// k: 剩余数字要分成k组
// used: 记录每个下标的数字是否已被使用
func dfsCanPartitionKSubsets(nums []int, start int, targetSum, currentSum int, k int, used []bool) bool {
	if k == 1 {
		// 若当前剩余元素只需要分成最后一组了，直接返回true即可
		return true
	} else if currentSum == targetSum {
		// 若当前已经找到一组，则继续寻找下一组 -- 也就是将下一个最大元素归位
		return dfsCanPartitionKSubsets(nums, len(nums)-1, targetSum, 0, k-1, used)
	}

	for i := start; i >= 0; i-- {
		// 若当前数字已被使用或当前数字相加后大于目标和 -- 剪枝
		if used[i] == true || currentSum+nums[i] > targetSum {
			continue
		}

		// 标记当前数字被使用
		used[i] = true

		// 从剩余数字中查找当前组合能否组成目标和，若能则返回true，若不能则说明当前数字不符合要求，将其移出选择队列，重新选择其他数字进行相加回溯
		if dfsCanPartitionKSubsets(nums, i-1, targetSum, currentSum+nums[i], k, used) {
			return true
		}

		// 当前数字未能组成目标和 -- 标记当前数字未被使用
		used[i] = false

		// 跳过重复元素 -- 剪枝
		for i > 0 && nums[i-1] == nums[i] {
			i--
		}
	}

	return false
}

// quickSort698 .
func quickSort698(nums []int, left, right int) {
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

	nums[l], nums[left] = nums[left], nums[l]
	quickSort698(nums, left, l-1)
	quickSort698(nums, r+1, right)
}
