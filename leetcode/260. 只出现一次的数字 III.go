package main

// 260. 只出现一次的数字 III

// 给定一个整数数组nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。
// 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
//
// 进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？

// func main() {

// }

// 类似offer 56、leetcode 136
func singleNumber(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	// 1. 所有元素求异或 -- 获得两个只出现一次的元素的异或值，找到分组依据
	flag := 0
	for _, n := range nums {
		flag ^= n
	}

	// 2. 将flag与自身的负数求与操作，找到第一个为1的位 -- 即找到两个数字的第一个不相同的位
	flag &= -flag

	// 3. 按照找到的分组依据将所有数字分成两组，分别组内求异或 -- 目的是将那两个只出现一次的数字分到两个组中
	res := make([]int, 2)
	for _, n := range nums {
		if n&flag == 0 {
			res[0] ^= n
		} else {
			res[1] ^= n
		}
	}

	return res
}
