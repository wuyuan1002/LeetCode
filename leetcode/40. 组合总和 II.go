package main

import "fmt"

// 40. 组合总和 II

// 给定一个数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
// candidates中的每个数字在每个组合中只能使用一次。
//
// 注意：解集不能包含重复的组合。

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum3([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

// 1. 回溯法
// 2. 动态规划 -- 01背包问题 -- 此题有重复元素，计算时无法去重
func combinationSum2(candidates []int, target int) [][]int {
	if candidates == nil || len(candidates) == 0 {
		return nil
	}

	sort1(candidates, 0, len(candidates)-1)

	result := make([][]int, 0)
	res := make([]int, 0)
	dfs7(candidates, 0, 0, target, &res, &result)

	return result
}

func dfs7(nums []int, start, currentSum, target int, res *[]int, result *[][]int) {
	if currentSum == target {
		tmp := make([]int, len(*res))
		copy(tmp, *res)
		*result = append(*result, tmp)
		return
	}

	for i := start; i < len(nums); i++ {
		if currentSum+nums[i] > target {
			return
		}

		// 跳过重复元素
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		*res = append(*res, nums[i])
		dfs7(nums, i+1, currentSum+nums[i], target, res, result)
		*res = (*res)[:len(*res)-1]
	}
}

func sort1(nums []int, left, right int) {
	if nums == nil || len(nums) == 0 || left >= right {
		return
	}

	l, r := left, right
	tmp := nums[left]
	for l < r {
		for l < r && nums[r] >= tmp {
			r--
		}
		for l < r && nums[l] <= tmp {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[l], nums[left] = tmp, nums[l]
	sort1(nums, left, l-1)
	sort1(nums, r+1, right)
}

func combinationSum3(candidates []int, target int) int {
	if candidates == nil || len(candidates) == 0 {
		return 0
	}

	// dp[i] += dp[i-num]
	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range candidates {
		for i := target; i >= 0; i-- {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}
