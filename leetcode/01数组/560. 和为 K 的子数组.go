package main

// 560. 和为 K 的子数组

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
//
// 子数组是数组中元素的连续非空序列。

// subarraySum .
// leetcode 523. 连续的子数组和
//
// 前缀和 -- 从数组第0项到当前项的和 -- prefixSum[i]表示nums[0, i]的子数组的前缀和
// 题意：有几种l、r的组合，使得nums[l, r]子数组区间的和等于k
//
// ↓ ↓ ↓ 转化为 ↓ ↓ ↓
//
// 有几种l、r的组合，满足prefixSum[r] - prefixSum[l-1] = k  ->  prefixSum[i-1] = prefixSum[j] - k
// 遍历数组的每个数字r，求出其前缀和sum并存入prefixSum，同时查看prefixSum中是否存在前缀和为sum-k的下标l，若存在则说明找到一个和为k的子数组
func subarraySum(nums []int, k int) int {
	// count：和为k的子数组个数、sum：nums从第0项到当前项的和
	count, sum := 0, 0
	// prefixSum[sum] = n：表示从第0项开始，前缀和为sum出现的次数为n，初始定义一个和为0出现的次数为1次 -- 用来快速查找一个前缀和出现的次数
	prefixSum := map[int]int{0: 1}

	// 不断遍历数组中的每一项，计算其前缀和sum，并统计map中有几项前缀和为sum-k的，更新出现次数并递增当前前缀和出现的次数
	for i := 0; i < len(nums); i++ {
		// 更新当前以当前数字结尾的前缀和
		sum += nums[i]

		// 查找在数组中下标nums[0, i)范围内前缀和为sum-k的出现过几次
		if _, ok := prefixSum[sum-k]; ok {
			count += prefixSum[sum-k]
		}

		// 将当前前缀和出现的次数+1
		prefixSum[sum] += 1
	}
	return count
}
