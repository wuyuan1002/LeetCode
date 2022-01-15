package main

// 268. 丢失的数字

// 给定一个包含 [0, n]中n个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
// 进阶：
// 你能否实现线性时间复杂度、仅使用额外常数空间的算法解决此问题?

// func main() {
// 	fmt.Println(missingNumber([]int{0, 3, 1}))
// }

// 1. 排序后判断
// 2. 放入map后判断
// 3. 先求出0-n的和，再求出数组中的和，两者做差
// 4，异或 -- 先求出[0,n]的异或值，再将该异或值与数组中的值挨个异或,
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
