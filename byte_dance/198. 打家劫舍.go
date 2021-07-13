package main

// 198. 打家劫舍

// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
// 影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

func main() {

}

// 动态规划
// 偷到某个房屋的最大价值 = max(当前房屋价值+前两个房屋最大价值, 前一个房屋的最大价值)
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	pre2 := nums[0]               // 前两个房屋
	pre1 := max(nums[0], nums[1]) // 前一个房屋
	for i := 2; i < len(nums); i++ {
		// 偷当前房屋的最大价值
		price := max(nums[i]+pre2, pre1)

		pre2 = pre1
		pre1 = price
	}

	return pre1
}
