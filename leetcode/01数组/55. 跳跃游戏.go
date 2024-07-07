package main

// 55. 跳跃游戏

// 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个下标。

// canJump .
// leetcode 45. 跳跃游戏 II
//
// 从某一位置起跳，能跳到的最远位置之前的位置都是可以跳到的，
// 所以，遍历每一个位置，不断更新能跳到的最远位置，过程中若发
// 现当前位置比最远位置都要远了，则说明跳不到数组最后一个下标
func canJump(nums []int) bool {
	farthest := 0 // 能跳到的最远距离

	// 不断向前移动，更新能跳到的最远位置
	for i := 0; i < len(nums); i++ {
		// 若当前位置已经超过能跳到的最远位置，说明跳不到数组最后
		if i > farthest {
			return false
		}

		// 更新能跳到的最远位置
		farthest = max(i+nums[i], farthest)
	}

	return true
}
