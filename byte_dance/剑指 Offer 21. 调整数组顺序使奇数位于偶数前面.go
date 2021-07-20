package main

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面

// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
// 使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

func main() {

}

// 双指针法
func exchange(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	// 判断是不是奇数
	isEven := func(num int) bool {
		return num&0x1 == 1
	}

	i, j := 0, len(nums)-1
	for i < j {
		for ; isEven(nums[i]) && i < j; i++ {
		}
		for ; !isEven(nums[j]) && i < j; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
