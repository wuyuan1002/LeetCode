package main

// 213. 打家劫舍 II

// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。
// 这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
// 同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，
// 今晚能够偷窃到的最高金额。

// func main() {

// }

// 同 198、337
// 房屋首尾相连，偷了第一个就不能偷最后一个，可以把环形房屋拆成两个单排房屋分别求最大值

// 环状排列意味着第一个房子和最后一个房子中只能选择一个偷窃，
// 因此可以把此环状排列房间问题约化为两个单排排列房间子问题：
//
// 在不偷窃第一个房子的情况下（即 nums[1:]），最大金额是 p1
// 在不偷窃最后一个房子的情况下（即 nums[:n-1]），最大金额是 p2
// 综合偷窃最大金额： 为以上两种情况的较大值，即 max(p1,p2)max(p1,p2) 。
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	robFront := robb(nums[:len(nums)-1]) // 第一个房屋可以偷的最大金额
	robTail := robb(nums[1:])            // 最后一个房屋可以偷的最大金额
	return max(robFront, robTail)
}

func robb(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	pre1, pre2 := max(nums[0], nums[1]), nums[0] // 前一个和前两个
	for i := 2; i < len(nums); i++ {
		price := max(nums[i]+pre2, pre1)
		pre2 = pre1
		pre1 = price
	}

	return pre1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
