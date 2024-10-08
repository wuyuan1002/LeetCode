package main

import "sort"

// 2294. 划分数组使最大差为 K

// 给你一个整数数组 nums 和一个整数 k 。你可以将 nums 划分成一个或多个 子序列 ，使 nums 中的每个元素都 恰好 出现在一个子序列中。
//
// 在满足每个子序列中最大值和最小值之间的差值最多为 k 的前提下，返回需要划分的 最少 子序列数目。
//
// 子序列 本质是一个序列，可以通过删除另一个序列中的某些元素（或者不删除）但不改变剩下元素的顺序得到。

// partitionArray .
// 先将数组从小到大进行排序，然后从最小的数字开始选择加入第一个子序列，直到该子序列中的数字差值大于k时，
// 说明该子序列为已满，然后从下一个数字开始加入下一个子序列，直到整个数组的数字都加入到子序列中
func partitionArray(nums []int, k int) int {
	// 总结果 -- 默认加入到第一个子序列，所以子序列的个数至少有一个
	result := 1

	// 先将数组从小到大进行排序
	sort.Ints(nums)

	// 不断从小到大选择数字加入到当前子序列中，若发现当前子序列差值大于k时，说明当前子序列已满，开始下一个子序列
	start := nums[0] // 当前子序列的首个元素
	for _, num := range nums {
		// 若当前子序列已满，则开始下一个子序列
		if num-start > k {
			result++
			start = num
		}
	}

	return result
}
