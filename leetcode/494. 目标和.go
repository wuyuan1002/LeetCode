package main

// 494. 目标和

func main() {

}

// 2. 动态规划 -- 背包问题 - 组合问题
// 同Hot100 518，类似Hot100 416
func findTargetSumWays(nums []int, S int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	// left组合 - right组合 = target
	// left + right等于sum，而sum是固定的
	// left - (sum - left) = target -> left = (target + sum)/2

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if S > sum || (sum+S)%2 == 1 {
		return 0
	}

	// 目标值
	target := (sum + S) / 2

	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range nums {
		for i := target; i >= 0; i-- {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
