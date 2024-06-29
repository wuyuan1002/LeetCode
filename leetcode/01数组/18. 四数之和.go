package main

import "sort"

// 18. 四数之和

// 给定一个包含n 个整数的数组nums和一个目标值target，
// 判断nums中是否存在四个元素 a，b，c和 d，使得a + b + c + d的值与target相等？
// 找出所有满足条件且不重复的四元组。
//
// 注意：答案中不可以包含重复的四元组。

// fourSum .
// 1. 动态规划，01背包问题 -- 可以求出解的个数
// 2. 回溯法 -- 可以求出各个解是什么
// 3. 类似三数之和，双指针 -- 比三数之和多固定一个数，多一层for循环
func fourSum(nums []int, target int) [][]int {
	if nums == nil || len(nums) < 4 {
		return nil
	}

	// 总结果
	result := make([][]int, 0)

	// 先将数组从小到大排序
	sort.Ints(nums)

	// 进行四数之和
	for i := 0; i < len(nums); i++ { // 固定第一个数
		if nums[i] > target && nums[i] >= 0 {
			// 剪枝
			// 若当前数字比target大且为正数，说明target大于0，说明当前数字后面的数字全是正数，相加只会比target更大 -- （正数+正数会越来越大）
			// 若当前数字比target大但为负数，说明target小于0，当前数字后面如果有负数的话四数之和可能会变小，因此不能剪枝负数的情况 -- （负数+负数会越来越小）
			break
		} else if i > 0 && nums[i] == nums[i-1] {
			// 剪枝 -- 跳过重复元素
			continue
		}

		// 进行三数之和
		for j := i + 1; j < len(nums); j++ { // 固定第二个数
			if nums[i]+nums[j] > target && nums[j] >= 0 {
				// 剪枝
				// 若当前两个数的和已经大于target且当前数字为正数，说明当前数字之后的数字只会越加越大，一定大于target了
				// 同理不能裁剪当前数字为负数的情况
				break
			} else if j > i+1 && nums[j] == nums[j-1] {
				// 剪枝 -- 跳过重复元素
				continue
			}

			// 双指针 -- 遍历剩余数字进行求解四数之和
			l, r := j+1, len(nums)-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum > target {
					// 右指针移动到下一个不相同数字位置
					for r--; l < r && nums[r] == nums[r+1]; r-- {
					}
				} else if sum < target {
					// 左指针移动到下一个不相同数字位置
					for l++; l < r && nums[l] == nums[l-1]; l++ {
					}
				} else {
					// 记录总结果 -- 四数之和等于tarhet
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})

					// 剪枝 -- 跳过重复元素，左右指针都移动到下一个不同数字位置
					for l++; l < r && nums[l] == nums[l-1]; l++ {
					}
					for r--; l < r && nums[r] == nums[r+1]; r-- {
					}
				}
			}
		}
	}

	return result
}
