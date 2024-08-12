package main

// LCR 120. 寻找文件副本

// 设备中存有 n 个文件，文件 id 记于数组 nums。
// 若文件 id 相同，则定义为该文件存在副本。请返回任一存在副本的文件 id。
//
// 提示：
// 0 ≤ nums[i] ≤ n-1
// 2 <= n <= 100000

// findRepeatDocument .
// Offer 03. 数组中重复的数字
//
// 原地交换
// 使用数组下标作为索引，将数字n交换到nums[n]处，若在交换时发现nums[n]已经为n，则说明数字n重复
func findRepeatDocument(nums []int) int {
	i, n := 0, 0 // 当前下标和对应的数字
	for i < len(nums) {
		// 获取当前数字
		n = nums[i]

		// 若当前数字和下标相等，说明它本身就在正确的位置
		if n == i {
			i++
			continue
		}

		// 若目标下标的值已经在正确位置，说明当前值是重复数字
		if nums[n] == n {
			return n
		}

		// 否则交换当前数字到正确的位置
		nums[i], nums[n] = nums[n], nums[i]
	}

	return -1
}
