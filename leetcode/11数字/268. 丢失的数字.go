package main

// 268. 丢失的数字

// 给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

// missingNumber .
// 先求出[0,n]的异或值，再将该异或值与数组中的值挨个异或，最终的异或值就是缺失的那个数字，因为其它的数字都异或了两遍
func missingNumber(nums []int) int {
	yh := 0
	for i := 0; i <= len(nums); i++ {
		yh ^= i
	}
	for i := 0; i < len(nums); i++ {
		yh ^= nums[i]
	}
	return yh
}
