package main

// 523. 连续的子数组和

// 给你一个整数数组 nums 和一个整数 k ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：
//
// 子数组大小 至少为 2 ，且
// 子数组元素总和为 k 的倍数。
// 如果存在，返回 true ；否则，返回 false 。
//
// 如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。0 始终视为 k 的一个倍数。

// checkSubarraySum .
// 同 leetcode 560. 和为 K 的子数组
// 前缀和 -- 从数组第0项到当前项的和 -- 相当于求数组中和为K的倍数的子数组
// sum[j] − sum[i−1] = n * k
// ↓ ↓ ↓ 转化为 ↓ ↓ ↓
// sum[j] / k − sum[i−1] / k = n  ->  要使得两者除k相减为整数，需要满足 sum[j] 和 sum[i−1] 对k取余相同
// 也就是说，我们只需要枚举右端点j，然后在枚举右端点j的时候检查[0, j)是否出现过左端点i，使得sum[j]和sum[i−1]对k取余相同
func checkSubarraySum(nums []int, k int) bool {
	// prefixSum[sum]: 从第0项开始，前缀和为sum的下标，存的是前缀和与k取余后的得数，因为这里要求的是对k取余相同的前缀和的下标 -- 初始定义一个和为0的下标为-1
	prefixSum := map[int]int{0: -1}

	// 不断遍历数组中的每一项，计算其前缀和sum，并查看map中前缀和为sum-k的下标是否满足条件
	sum := 0 // sum: nums从第0项到当前项的和
	for i := 0; i < len(nums); i++ {
		// 更新当前以当前数字结尾的前缀和
		sum += nums[i]

		// 判断map中前缀和为 sum%k 是否出现过，若出现过且子数组长度大于2则说明存在满足条件的子数组
		// prefixSum存的是前缀和与k取余后的得数，因为这里要求的是对k取余相同的前缀和的下标
		if idx, ok := prefixSum[sum%k]; ok {
			if i-idx >= 2 {
				return true
			}
		} else {
			// 标记当前前缀和出现的下标 -- 若当前前缀和已经出现过，则不会进行更新，只记录某个前缀和第一次出现时的下标
			prefixSum[sum%k] = i
		}

	}
	return false
}
