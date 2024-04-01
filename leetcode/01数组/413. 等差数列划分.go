package main

// 413. 等差数列划分

// 如果一个数列 至少有三个元素 ，并且任意两个相邻元素之差相同，则称该数列为等差数列。
//
// 例如，[1,3,5,7,9]、[7,7,7,7] 和 [3,-1,-5,-9] 都是等差数列。
// 给你一个整数数组 nums ，返回数组 nums 中所有为等差数组的 子数组 个数。
//
// 子数组 是数组中的一个连续序列。

// numberOfArithmeticSlices .
// 1. 首先遍历数组求出每个相邻元素之间的差存在一个数组a中，然后多次遍历数组a计算长度为3、4、5...的子数组个数
// 2. 差分+计数 -- 使用一个变量保存当前差值，遍历一次数组，计算与前一个数字的差，若差相同则说明是等差数列，累加总结果
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	result := 0

	// 当前数字与前一个数字的差、当前差出现的次数
	diff, count := nums[1]-nums[0], 0

	// 遍历每个数字计算等差数列个数 -- 因为等差数列的长度至少为3，所以可以从第3个数字开始遍历
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == diff {
			count++
		} else {
			diff, count = nums[i]-nums[i-1], 0
		}

		// 递增计算总结果
		result += count
	}

	return result
}
