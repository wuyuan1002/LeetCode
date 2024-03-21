package main

import "sort"

// 1005. K 次取反后最大化的数组和

// 给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
//
// 选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
// 重复这个过程恰好 k 次。可以多次选择同一个下标 i 。
//
// 以这种方式修改数组后，返回数组 可能的最大和 。

// largestSumAfterKNegations .
// 贪心的思路，局部最优：让绝对值大的负数变为正数，当前数值达到最大，整体最优：整个数组和达到最大
// 第一步：将数组按照绝对值大小从大到小排序
// 第二步：从前向后遍历，遇到负数将其变为正数，同时K--
// 第三步：如果K还大于0，那么反复转变数值最小的元素，将K用完
// 第四步：求和
func largestSumAfterKNegations(nums []int, k int) int {

	// 将数组根据绝对值从大到小排序
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) > abs(nums[j])
	})

	// 从前向后遍历数组，k>0时将负数取反
	for i := range nums {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
	}

	// 如果K还大于0，那么反复转变数值最小的元素，将K用完
	// 当k与2取余为1时，说明k是奇数，需要将最小元素变为负数，否则说明k为偶数，负负得正了
	if k%2 == 1 {
		nums[len(nums)-1] *= -1
	}

	// 累加数组元素
	result := 0
	for _, n := range nums {
		result += n
	}

	return result
}
