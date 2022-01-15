package main

// 53. 最大子序和

// func main() {

// }

func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	sum := nums[0]
	maxSum := sum
	for i := 1; i < len(nums); i++ {
		sum = max(sum+nums[i], nums[i])
		maxSum = max(maxSum, sum)
	}

	return maxSum
}
