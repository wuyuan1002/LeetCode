package main

// 485. 最大连续 1 的个数

// 给定一个二进制数组， 计算其中最大连续 1 的个数。

func main() {

}

func findMaxConsecutiveOnes(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	maxLen := 0 // 最大长度
	lenth := 0  // 当前长度
	for _, n := range nums {
		if n == 1 {
			lenth++
		} else if n == 0 {
			lenth = 0
		}
		maxLen = max(maxLen, lenth)
	}

	return maxLen
}
