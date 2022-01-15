package main

// 18. 四数之和

// 给定一个包含n 个整数的数组nums和一个目标值target，
// 判断nums中是否存在四个元素 a，b，c和 d，使得a + b + c + d的值与target相等？
// 找出所有满足条件且不重复的四元组。
//
// 注意：答案中不可以包含重复的四元组。

// func main() {

// }

// 1. 动态规划，01背包问题 -- 可以求出解的个数
// 2. 回溯法 -- 可以求出各个解是什么
// 3. 类似三数之和，双指针 -- 比三数之和多固定一个数，多一层for循环
func fourSum(nums []int, target int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	quickSort(nums, 0, len(nums)-1)

	res := make([][]int, 0)

	for i := 0; i < len(nums); i++ { // 先固定一个数
		if i > 0 && nums[i] == nums[i-1] {
			// 跳过重复元素
			continue
		}

		for j := i + 1; j < len(nums); j++ { // 再固定一个数
			if j > i+1 && nums[j] == nums[j-1] {
				// 跳过重复元素
				continue
			}

			l, r := j+1, len(nums)-1 // 在之后的元素中使用双指针
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})

					for l++; l < r && nums[l] == nums[l-1]; l++ {
					}
					for r--; l < r && nums[r] == nums[r+1]; r-- {
					}
				}
			}
		}
	}

	return res
}
