package main

// 560. 和为K的子数组

// 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

// func main() {
// 	fmt.Println(subarraySum([]int{1, 2, -2, 4, 3, 2}, 4))
// }

// 前缀和
// 题意：有几种 i、j 的组合，使得从第 i 到 j 项的子数组和等于 k。
// ↓ ↓ ↓ 转化为 ↓ ↓ ↓
// 有几种 i、j 的组合，满足 prefixSum[j] - prefixSum[i - 1] == kprefixSum[j]−prefixSum[i−1]==k。
func subarraySum(nums []int, k int) int {
	count, preSum := 0, 0          // preSum: nums 的第 0 项到 当前项 的和
	prefixSum := make(map[int]int) // prefixSum[x]: 第 0 项到 第 x 项 的和
	prefixSum[0] = 1
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if _, ok := prefixSum[preSum-k]; ok {
			count += prefixSum[preSum-k]
		}
		prefixSum[preSum] += 1
	}
	return count
}
