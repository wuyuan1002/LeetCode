package main

import "math"

// 41. 缺失的第一个正数

// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

// firstMissingPositive .
// leetcode 442. 数组中重复的数据
// leetcode 448. 找到所有数组中消失的数字
//
// 原地哈希
// 使用 下标k 处数字的正负号表示 数字k+1 是否在数组中出现过，标记结束后，第一个值大于0的下标+1就是缺失的第一个正数
// 如数组中下标0处数字的正负值表示数字1是否出现过、下标1处数字的正负值表示数字2是否出现过
//
// 数组长度为n，那么第一个缺失的正数只能在[1, n]区间内或等于n+1
// - 在[1, n]区间内说明数组中存在小于0、大于n、或重复的数字 -- 数组空间一共就n个，还存在不合法或重复数字占用了空间，那么缺失的第一个合法数字一定在[1, n]内
// - 为n+1说明数组中数字全部为[1, n]区间内且没有重复数字 -- 数组所有空间存的都是合法数字
func firstMissingPositive(nums []int) int {
	// 遍历数组，将所有不合法数字（不在[1, n]范围内的数字）标记为MaxInt32
	// 遍历结束后，数组中所有数字都大于0，其中在[1, n]范围内的为合法数字，等于MaxInt32的为不合法数字
	for i, n := range nums {
		if n <= 0 || n > len(nums) {
			nums[i] = math.MaxInt32
		}
	}

	// 遍历数组，将数组中所有合法的数字的下标对应的数字标记为负数表示该数字已经出现过
	// 遍历结束后，数组中为负数的位置说明其下标对应数字出现过，正数的位置说明其下标对应的数字没出现过
	for i := 0; i < len(nums); i++ {
		// 获取当前数字的绝对值 -- 因为当前位置数字可能为负数, 所以要使用绝对值
		num := abs(nums[i])

		// 若当前数字为合法数字且对应位置还是正数（为正数说明当前数字是第一次出现，标记过一次后就不用再继续标记了，防止当前数字出现多次出现负负得正的情况）
		if num >= 1 && num <= len(nums) && nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}

	// 遍历数组，第一个数值大于0的下标+1就是缺失的第一个正数
	for i, n := range nums {
		if n > 0 {
			return i + 1
		}
	}

	// 数组中所有数字都小于0，说明数组中数字全部为[1, n]区间内且没有重复数字，缺失的第一个正数为n+1
	return len(nums) + 1
}

// abs .
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
