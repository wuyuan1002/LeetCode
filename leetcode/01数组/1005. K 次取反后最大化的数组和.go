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
// 优先将绝对值大的负数反转为正数，若负数的数量大于k，那么就不断反转数值最小的数字将k用完
func largestSumAfterKNegations(nums []int, k int) int {

	// 将数组根据绝对值从大到小排序
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) > abs(nums[j])
	})

	// 从前向后遍历数组，按照负数的绝对值从大到小进行反转
	for i := range nums {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
	}

	// 如果k还大于0，说明仅反转负数无法将k用完，那么就不断转变数值最小的元素，将k用完
	// - 若剩余的k为奇数，反转奇数次相当于反转1次，所以这里直接将最小的数字反转一次即可
	// - 若剩余的k为偶数，反转偶数次相当于没有反转，所以这里可以不做反转操作
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
