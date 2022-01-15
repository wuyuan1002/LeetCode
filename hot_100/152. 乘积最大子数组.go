package main

// 152. 乘积最大子数组

// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），
// 并返回该子数组所对应的乘积。

// func main() {
// 	fmt.Println(maxProduct([]int{-2, 3, -4}))
// }

// 动态规划
// 类似offer 42
func maxProduct(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	// 以某个值结尾的最大乘积 = max(当前元素, max(以上一个元素结尾的的最小乘积 * 当前元素, 以上一个元素结尾的的最大乘积 * 当前元素))
	// 以上一个元素结尾的子数组乘积的最小值、最大值，以及总结果的最大值
	minPre, maxPre, maxRes := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// 当前元素与前一个元素的最大最小值的乘积
		mn := minPre * nums[i]
		mx := maxPre * nums[i]

		// 计算与当前元素相乘后的最大最小值
		minP := min(nums[i], min(mn, mx))
		maxP := max(nums[i], max(mn, mx))

		// 更新记录的前一个元素的乘积最大最小值
		minPre = min(minP, maxP)
		maxPre = max(minP, maxP)

		// 更新总结果
		maxRes = max(maxRes, max(minP, maxP))
	}

	return maxRes
}
