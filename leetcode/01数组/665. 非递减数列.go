package main

// 665. 非递减数列

// 给你一个长度为 n 的整数数组 nums ，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。
//
// 我们是这样定义一个非递减数列的： 对于数组中任意的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。

// checkPossibility .
func checkPossibility(nums []int) bool {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		// 若当前元素比下一个元素大，说明找到了一个递减的顺序
		if nums[i] > nums[i+1] {
			// 递增递减序列次数，若递减次数大于1说明不可能只修改一个数字使整体递增
			count++
			if count > 1 {
				return false
			}

			// 确保 nums[i−1] <= nums[i+1]
			if i > 0 && nums[i+1] < nums[i-1] {
				nums[i+1] = nums[i]
			}
		}
	}
	return true
}
