package main

// 260. 只出现一次的数字 III

// 给你一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。
// 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
//
// 你必须设计并实现线性时间复杂度的算法且仅使用常量额外空间来解决此问题。

// singleNumber260 .
// 数组中有两个数字只出现一次，其余数字出现了两次，我们只需要想办法将这两个只出现一次的数字划分到两个组，然后对每个组进行136的求法找出这个只出现一次的数字即可
// 1. 将所有数字进行异或，因为出现两次的数字异或后为0，所以最终的异或结果就是两个只出现一次的数字的异或结果flag
// 2. flag右侧第一个二进制为1的位即为两个只出现一次的数字的第一个不同的位，我们使用这个位作为分组的依据 -- 将flag与自身的负数求与操作，即可找到右侧第一个为1的位
// 3. 按照找到的分组依据将所有数字分成两组，分别组内求异或 -- 目的是将那两个只出现一次的数字分到两个组中
// 4. 返回每组中只出现一次的数字
func singleNumber260(nums []int) []int {
	// 1. 所有元素求异或 -- 获得两个只出现一次的元素的异或值，找到分组依据
	flag := 0
	for _, n := range nums {
		flag ^= n
	}

	// 2. 将flag与自身的负数求与操作，找到其右侧第一个为1的位 -- 即找到两个数字的右侧第一个不相同的位
	flag &= -flag

	// 3. 按照flag作为分组依据将所有数字分成两组，分别组内求异或 -- 目的是将那两个只出现一次的数字分到两个组中
	result := make([]int, 2)
	for _, num := range nums {
		if num&flag == 0 {
			result[0] ^= num
		} else {
			result[1] ^= num
		}
	}

	// 4. 返回每组中只出现一次的数字
	return result
}
