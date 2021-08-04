package main

// 416. 分割等和子集

func main() {

}

// Hot100 322, 518
// 背包问题 - 存在问题 -- dp[i]=dp[i] || dp[i-num]
func canPartition(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2            // 目标和
	dp := make([]bool, target+1) // dp[i] = dp[i] || dp[i-num]
	dp[0] = true
	for _, n := range nums {
		for i := target; i >= 0; i-- {
			if i >= n {
				dp[i] = dp[i] || dp[i-n]
			}
		}
	}
	return dp[target]
}
