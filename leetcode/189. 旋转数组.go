package main

// 189. 旋转数组

// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//
// 进阶：
// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
// 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

// 类似Offer 58、leetcode 151
// 类似 61. 旋转链表
func rotate(nums []int, k int) {
	if k <= 0 {
		return
	}

	realK := k % len(nums)
	if realK == 0 {
		return
	}

	reverse1(nums)
	reverse1(nums[:realK])
	reverse1(nums[realK:])
}

func reverse1(nums []int) {
	if len(nums) == 0 {
		return
	}

	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
