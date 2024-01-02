package main

// 136. 只出现一次的数字

// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
// 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。

// singleNumber .
// 异或（相同为0不同为1），所以一个数字与其自身异或后为0，因为数组中除一个元素只出现一次，其余都出现两次，
// 所以将数组中所有数字进行异或，出现两次的数字都为0，结果就是只出现一次的数字与0进行异或，即那个只出现一次的数字
func singleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}
