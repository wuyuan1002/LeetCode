package main

import (
	"math"
	"strconv"
)

// 556. 下一个更大元素 III

// 给你一个正整数 n ，请你找出符合条件的最小整数，其由重新排列 n 中存在的每位数字组成，
// 并且其值大于 n 。如果不存在这样的正整数，则返回 -1 。
//
// 注意 ，返回的整数应当是一个 32 位整数 ，如果存在满足题意的答案，
// 但不是 32 位整数 ，同样返回 -1 。

// nextGreaterElement556 .
// 同 leetcode 31. 下一个排列
// 本题其实是在求大于n的最小的整数x，但是x只能使用n中的每一位数字
// 使用一个数组存放n的每一位，然后求该数组的下一个排列
func nextGreaterElement556(n int) int {
	// 1. 先把n转换成数字切片，对应的其实是每一位数字的ascii码
	nums := []byte(strconv.Itoa(n))

	// 求nums的下一个排列

	// 2. 找到满足比它的下一个元素小的最大的索引k
	k := len(nums) - 2
	for ; k >= 0; k-- {
		if nums[k] < nums[k+1] {
			break
		}
	}

	// 若1.不存在，说明数组已经是倒序排列了，最大，题目要求返回-1
	if k < 0 {
		return -1
	}

	// 3. 找出满足比nums[k]大的最大索引l
	l := len(nums) - 1
	for ; l >= 0; l-- {
		if nums[l] > nums[k] {
			break
		}
	}

	// 4. 交换 nums[l] 和 nums[k]
	nums[k], nums[l] = nums[l], nums[k]

	// 5. 反转nums[k+1:]
	reverse(nums[k+1:])

	// 返回结果 -- 题目要求最大为32位整数，若不是32位也返回-1
	result, _ := strconv.Atoi(string(nums))
	if result > math.MaxInt32 {
		result = -1
	}
	return result
}

// reverse 反转数组
func reverse(nums []byte) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
