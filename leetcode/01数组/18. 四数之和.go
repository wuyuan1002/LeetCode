package main

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

	quickSort(nums, 0, len(nums)-1)

	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ { // 先固定一个数
		if nums[i] > target && nums[i] >= 0 {
			break
		} else if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ { // 进行三数之和
			if nums[i]+nums[j] > target && nums[j] >= 0 {
				break
			} else if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, len(nums)-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum > target {
					// 移动到下一个不相同数字位置
					for r--; l < r && nums[r] == nums[r+1]; r-- {
					}
				} else if sum < target {
					// 移动到下一个不相同数字位置
					for l++; l < r && nums[l] == nums[l-1]; l++ {
					}
				} else {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})

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
