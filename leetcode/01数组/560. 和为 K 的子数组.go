package main

// 560. 和为 K 的子数组

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
//
// 子数组是数组中元素的连续非空序列。

// subarraySum .
// 同 leetcode 523. 连续的子数组和
// 前缀和 -- 从数组第0项到当前项的和
// 题意：有几种 i、j 的组合，使得从第 i 到 j 项的子数组和等于 k
// ↓ ↓ ↓ 转化为 ↓ ↓ ↓
// 有几种 i、j 的组合，满足 prefixSum[j] - prefixSum[i-1] = k  ->  prefixSum[i-1] = prefixSum[j] - k
// 可以通过求出 prefixSum 数组的每一项，再看哪些项相减等于k，求出count
func subarraySum(nums []int, k int) int {
	count, sum := 0, 0             // sum: nums从第0项到当前项的和
	prefixSum := map[int]int{0: 1} // prefixSum[sum]: 从第0项开始，前缀和为sum的次数 -- 初始定义一个和为0出现的次数为1次

	// 不断遍历数组中的每一项，计算其前缀和sum，并统计map中有几项前缀和为sum-k的，更新出现次数并递增当前前缀和出现的次数
	for i := 0; i < len(nums); i++ {
		// 更新当前以当前数字结尾的前缀和
		sum += nums[i]

		// 查找在数组中下标[0, i)范围内前缀和为 sum-k 的出现过几次
		if _, ok := prefixSum[sum-k]; ok {
			count += prefixSum[sum-k]
		}

		// 将当前前缀和出现的次数+1
		prefixSum[sum] += 1
	}
	return count
}
