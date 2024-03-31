package main

// 1567. 乘积为正数的最长子数组长度

// 给你一个整数数组 nums ，请你求出乘积为正数的最长子数组的长度。
//
// 一个数组的子数组是由原数组中零个或者更多个连续数字组成的数组。
//
// 请你返回乘积为正数的最长子数组长度。

// getMaxLen .
// 同 leetcode 376. 摆动序列
//
// positive[i]表示以第i个数字结尾的乘积为正数的最长子数组长度
// negative[i]表示以第i个数字结尾的乘积为负数的最长子数组长度
func getMaxLen(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 构造两个dp数组，分别表示乘积为正数的最长子数组长度和乘积为负数的最长子数组长度
	positive, negative := make([]int, len(nums)), make([]int, len(nums))

	// 初始化dp数组 -- 根据数组中的第一个数字是正数还是负数初始化dp数组的第一个位置
	if nums[0] > 0 {
		positive[0] = 1
	} else if nums[0] < 0 {
		negative[0] = 1
	}

	// 总结果 -- 乘积为正数的最长子数组长度，初始值为第一个数字是否为正数，如[1, 0, 0, 0, 0, 0]
	maxLengh := positive[0]

	// 从第二个数字开始遍历每一个数字，更新以当前数字结尾的乘积为正数和负数的最大子数组长度，并计算乘积为正数的最长子数组长度
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			// 若当前数字大于0，则乘积为正数的子数组长度直接+1，乘积为负数的子数组长度若原来没有负数则保持为0否则+1
			positive[i] = positive[i-1] + 1
			if negative[i-1] > 0 {
				negative[i] = negative[i-1] + 1
			}
		} else if nums[i] < 0 {
			// 若当前数字小于0，则乘积为负数的子数组长度为乘积为正数的子数组长度+1，乘积为负数的子数组长度由乘积为正数的子数组推倒而来
			negative[i] = positive[i-1] + 1
			if negative[i-1] > 0 {
				positive[i] = negative[i-1] + 1
			}
		} else {
			// 若当前数字为0，则任何数字乘0都为0，子数组长度置为0重新计算
			positive[i] = 0
			negative[i] = 0
		}

		// 更新乘积为正数的最大子数组长度
		maxLengh = max(maxLengh, positive[i])
	}

	return maxLengh
}

// getMaxLen1 滚动数组优化
// 可以看出在求positive[i]和negative[i]时，只依赖其前一个数字的结果，即positive[i-1]和negative[i-1]，
// 因此，positive和和negative可以只使用一个变量用来记录前一个下标所能构成的最长递增摆动子序列和最长递减摆动子序列的长度即可，
// 而不必使用数组将每一个下标的都记录下来
func getMaxLen1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	positive, negative := 0, 0

	if nums[0] > 0 {
		positive = 1
	} else if nums[0] < 0 {
		negative = 1
	}

	maxLengh := positive

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			positive++
			if negative > 0 {
				negative = negative + 1
			}
		} else if nums[i] < 0 {
			newPositive := 0
			if negative > 0 {
				newPositive = negative + 1
			}
			newNegative := positive + 1

			positive = newPositive
			negative = newNegative
		} else {
			positive = 0
			negative = 0
		}

		maxLengh = max(maxLengh, positive)
	}

	return maxLengh
}
