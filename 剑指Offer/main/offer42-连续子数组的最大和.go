package main

// 剑指 Offer 42. 连续子数组的最大和

// 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
// 要求时间复杂度为O(n)。

// func main() {

// }

// 类似Hot100 152
// 使用动态规划 -- 见offer14
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		panic("nums is empty")
	}

	// 以某个值结尾的最大和 = max(它本身+以前一个值结尾的最大和, 它本身)
	products := make([]int, len(nums))
	products[0] = nums[0] // products[i]表示以第i个数结尾的最大和

	sum := nums[0] // 最大和
	for i := 1; i < len(nums); i++ {
		products[i] = max(nums[i]+products[i-1], nums[i])
		sum = max(products[i], sum)
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 第二次做 -- 不用动态规划，每次加数组后面一个元素，前面的和小于当前数字时抛弃前面的数字重新累加
func maxSubArray1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	maxSub := nums[0]     // 记录的最大和
	currentSub := nums[0] // 当前的和
	for i := 1; i < len(nums); i++ {
		currentSub += nums[i]     // 先把数字加到当前和上
		if currentSub < nums[i] { // 若加上这个数字后比这个数字本身还小，则舍弃前面的数字，从当前数字开始继续象后累加
			currentSub = nums[i]
		}

		maxSub = max(maxSub, currentSub) // 计算最大和
	}
	return maxSub
}
