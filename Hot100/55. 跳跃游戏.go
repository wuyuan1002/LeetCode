package main

// 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个下标。

func main() {

}

func canJump(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return true
	}
	return dfs4(nums, 0)
}

// 回溯法 -- 超出时间限制
// index: 当前所在的数组下标
func dfs4(nums []int, index int) bool {
	// 正好跳到最后一个下标
	if index == len(nums)-1 {
		return true
	}

	for i := 1; i <= nums[index]; i++ {
		if dfs4(nums, index+i) {
			return true
		}
	}

	return false
}
