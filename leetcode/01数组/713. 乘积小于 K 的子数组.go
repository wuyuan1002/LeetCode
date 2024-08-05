package main

// 713. 乘积小于 K 的子数组

// 给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。

// numSubarrayProductLessThanK .
// leetcode 209. 长度最小的子数组
//
// 滑动窗口
// 不断向右移动右边界，直到窗口中数字的乘积大于k后向右移动左边界缩小窗口，直到窗口总乘积小于k后继续向后移动右边界，
// 每一次向后移动右边界，都将由于这个新元素的加入所带来的新连续子数组的个数加到总结果
func numSubarrayProductLessThanK(nums []int, k int) int {
	// 总结果、当前滑动窗口内所有数字的乘积
	result, prod := 0, 1

	for l, r := 0, 0; r < len(nums); r++ {
		// 将当前新加入的右边界数字乘到当前窗口乘积中，表示将当前数字加入窗口
		prod *= nums[r]

		// 若窗口内数字乘积比k大了，则不断向前移动左边界直到窗口内乘积小于k
		for ; l <= r && prod >= k; l++ {
			prod /= nums[l]
		}

		// 将由于当前右边界加入窗口所带来的新连续子数组的个数加到总结果
		result += r - l + 1
	}

	return result
}
