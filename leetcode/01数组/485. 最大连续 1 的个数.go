package main

// 485. 最大连续 1 的个数

// 给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。

// findMaxConsecutiveOnes .
// 使用一个变量统计连续1的个数，若遇到0则归零重新统计
func findMaxConsecutiveOnes(nums []int) int {
	// 最大连续1的个数、当前连续1的个数
	result, count := 0, 0

	// 遍历数组中每个数字，不断统计最大连续1的个数并更新最大值
	for _, n := range nums {
		// 若当前数字为1则递增连续1的计数，若为0则将连续1的计数归零重新统计
		if n == 1 {
			count++
		} else {
			count = 0
		}

		// 更新最大连续1的个数
		result = max(result, count)
	}

	return result
}
