package main

// 1004. 最大连续1的个数 III

// 给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。

// longestOnes .
// 同 leetcode 424. 替换后的最长重复字符
// 滑动窗口 -- 求窗口中0的个数小于k的最大窗口长度
func longestOnes(nums []int, k int) int {

	result := 0

	// 窗口左端点、窗口中0的个数
	l, count := 0, 0

	// 不断向右移动右端点，维护窗口中0的个数最大为k个，求窗口的最大长度
	for r := 0; r < len(nums); r++ {
		// 若当前数字为0，则增加窗口中0的计数
		if nums[r] == 0 {
			count++
		}

		// 若窗口中0的个数大于k了，则向右移动左端点使窗口中0的个数符合要求
		if count > k {
			for count > k {
				if nums[l] == 0 {
					count--
				}
				l++
			}
		}

		// 计算窗口长度，统计最大满足条件的窗口长度
		result = max(result, r-l+1)
	}

	return result
}
