package main

import (
	"fmt"
	"math"
)

// 41. 缺失的第一个正数

// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

func main() {
	fmt.Println(firstMissingPositive([]int{1}))
}

// 1. 将所有数字放入map，随后从1开始依次枚举正整数，并判断其是否在哈希表中
// 2. 使用数组构建这个hsah表
func firstMissingPositive(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	// 先将不合法的数字标记出来
	for i, n := range nums {
		if n <= 0 || n > len(nums) {
			nums[i] = math.MaxInt32
		}
	}

	// 将合法数字数值对应的下标标记为对应的负数
	for _, n := range nums {
		num := int(math.Abs(float64(n)))
		if num <= len(nums) {
			nums[num-1] = -int(math.Abs(float64(nums[num-1])))
		}
	}

	// 遍历数组，第一个数值大于0的下标+1就是缺失的第一个正数
	for i, n := range nums {
		if n > 0 {
			return i + 1
		}
	}

	return len(nums) + 1
}
