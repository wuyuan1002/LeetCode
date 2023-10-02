package main

// 39. 组合总和

// func main() {

// }

// 1. 回溯法
// 2. 动态规划 -- 完全背包问题
func combinationSum(candidates []int, target int) [][]int {
	if candidates == nil || len(candidates) == 0 {
		return nil
	}

	sort2(candidates, 0, len(candidates)-1)

	result := make([][]int, 0)
	res := make([]int, 0)
	dfs6(candidates, 0, 0, target, &res, &result)

	return result
}

func dfs6(nums []int, start, currentSum, target int, res *[]int, result *[][]int) {
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

		*res = append(*res, nums[i])
		dfs6(nums, i, currentSum+nums[i], target, res, result)
		*res = (*res)[:len(*res)-1]
	}
}

func combinationSum1(candidates []int, target int) int {
	if candidates == nil || len(candidates) == 0 {
		return 0
	}

	// dp[i] += dp[i-num]
	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range candidates {
		for i := 1; i <= target; i++ {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}
