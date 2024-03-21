package main

// 152. 乘积最大子数组

// 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），
// 并返回该子数组所对应的乘积。

// maxProduct .
// 以某个值结尾的最大乘积 = max(当前元素, max(以前一个元素结尾的的最大乘积 * 当前元素, 以前一个元素结尾的的最小乘积 * 当前元素))
// 之所以不直接只用最大乘积是因为若当前数字是负数，前一个的最小乘积可以与当前数字相乘负负得正得到最大乘积
func maxProduct(nums []int) int {
	// 总结果的最大值
	result := nums[0]
	// 以前一个元素结尾的子数组乘积的最小值、最大值
	minPre, maxPre := nums[0], nums[0]

	// 遍历数组的所有数字，计算以每个数字结尾的最大和最小乘积，并更新总结果
	for i := 1; i < len(nums); i++ {
		// 使用两个变量记录以前一个数字结尾的最大最小乘积
		maxP, minP := maxPre, minPre

		// 计算以当前数字结尾的最大和最小乘积
		maxPre = max(nums[i], max(maxP*nums[i], minP*nums[i]))
		minPre = min(nums[i], min(minP*nums[i], maxP*nums[i]))

		// 更新总结果
		result = max(result, maxPre)
	}

	return result
}
