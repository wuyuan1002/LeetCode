package main

import (
	"math"
	"sort"
)

// 16. 最接近的三数之和

// 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。
// 请你从 nums 中选出三个整数，使它们的和与 target 最接近。
// 返回这三个数的和。
//
// 假定每组输入只存在恰好一个解。

// threeSumClosest .
// 同 leetcode 15. 三数之和
// 使用绝对值判断与目标值的距离
func threeSumClosest(nums []int, target int) int {
	result := math.MaxInt32

	// 先将数组从小到大排序
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			// 剪枝 -- 跳过重复元素
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			// 收缩双指针
			if sum > target {
				// 右指针移动到下一个不相同数字位置
				for r--; l < r && nums[r] == nums[r+1]; r-- {
				}
			} else if sum < target {
				// 左指针移动到下一个不相同数字位置
				for l++; l < r && nums[l] == nums[l-1]; l++ {
				}
			} else {
				// 若当前三数之和正好等于target则直接返回，因为不会有比target本身更接近target的值了
				return sum
			}

			// 若当前三数之和比原来结果更接近target则更新总结果
			if abs(sum-target) < abs(result-target) {
				result = sum
			}
		}
	}

	return result
}
