package main

// 45. 跳跃游戏 II

// 给你一个非负整数数组nums ，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
// 假设你总是可以到达数组的最后一个位置。

func main() {

}
func jump(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	step, end, far := 0, 0, 0 // 跳跃次数、本次step能跳到的最远位置、本次step的下一跳能跳到的最远位置
	for i := 0; i < len(nums)-1; i++ {
		far = max(far, i+nums[i])
		if i == end {
			end = far
			step++
		}
	}

	return step
}
