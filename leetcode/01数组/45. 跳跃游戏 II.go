package main

// 45. 跳跃游戏 II

// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
//
// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。
// 换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
//
// 0 <= j <= nums[i]
// i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

// jump .
// leetcode 55. 跳跃游戏
//
// 从某一位置起跳，能跳到的最远位置之前的位置都是可以跳到的，
// 遍历一次起跳到能跳到的最远位置之间，下次step能跳到的最远位置，争取让每次跳跃都跳到数字大的地方
func jump(nums []int) int {
	// 跳跃次数、本次step能跳到的最远位置、本次step的下一跳能跳到的最远位置
	step, end, farthest := 0, 0, 0

	// 不断遍历一次起跳到本次step能跳到的最远位置之间[i, end]，统计下一跳能跳到的最远位置，争取让每次跳跃都跳到数字大的地方
	for i := 0; i < len(nums)-1; i++ {
		// 更新下次step能跳到的最远距离
		farthest = max(i+nums[i], farthest)

		// 若本次step已经跳到它能跳到的最远位置，说明已经统计完本次step的下一次step能跳到的最远位置
		// 本次step结束，将end赋值为下一次step能跳到的最远位置，递增跳跃次数，开始下一次step
		if i == end {
			end = farthest
			step++
		}
	}

	return step
}
