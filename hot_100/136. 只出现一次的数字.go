package main

// 136. 只出现一次的数字

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
// 说明：
// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

// func main() {

// }

// offer 56， leetcode 260
// 异或运算 -- 一个数字与本身异或后为0
func singleNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	num := 0
	for _, n := range nums {
		num ^= n
	}

	return num
}
