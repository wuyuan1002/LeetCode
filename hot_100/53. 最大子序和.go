package main

// 53. 最大子序和

// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

// func main() {
// 	fmt.Println(maxSubArray([]int{-1, -2}))
// }

func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sum := nums[0]    // 当前和
	maxsum := nums[0] // 最大和

	for i := 1; i < len(nums); i++ {
		// 更新当前和
		sum = max(sum+nums[i], nums[i])
		// 更新最大和
		maxsum = max(sum, maxsum)
	}

	return maxsum
}
