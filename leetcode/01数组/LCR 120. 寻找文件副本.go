package main

// LCR 120. 寻找文件副本

// 设备中存有 n 个文件，文件 id 记于数组 documents。
// 若文件 id 相同，则定义为该文件存在副本。请返回任一存在副本的文件 id。
//
// 提示：
// 0 ≤ documents[i] ≤ n-1
// 2 <= n <= 100000

// findRepeatDocument .
// Offer 03. 数组中重复的数字
//
// 根据题目提示可知数组中数字的大小都小于数组长度，因此可以进行原地交换，使用数组下标作为索引，
// 因为每个数字都在0~n-1范围内，所以如果所有数字都不重复的话，每个数字对应一个下标，但是若有数字重复的话
// 就会有一个数字对应多个下标。因此，可以遍历数组，把遇到的每一个数字都放到它对应的下标处，如果在放某个数字时发现
// 下标处的数字已经和下标相等了，那么这个数字就是重复数字
func findRepeatDocument(documents []int) int {
	i, n := 0, 0 // 当前下标和对应的数字
	for i < len(documents) {
		// 获取当前数字
		n = documents[i]

		// 若当前数字和下标相等，说明它本身就在正确的位置
		if n == i {
			i++
			continue
		}

		// 若目标下标的值已经在正确位置，说明当前值是重复数字
		if documents[n] == n {
			return n
		}

		// 否则交换当前数字到正确的位置
		documents[i], documents[n] = documents[n], documents[i]
	}

	return -1
}
