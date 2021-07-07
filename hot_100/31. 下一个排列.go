package main

import (
	"fmt"
)

// 31. 下一个排列

// 实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
// 必须 原地 修改，只允许使用额外常数空间。

func main() {
	aa := []int{3, 1, 2}
	nextPermutation(aa)
	fmt.Println(aa)
}

// 按字典顺序生成序列：
// 1.先找出最大的索引 k 满足 nums[k] < nums[k+1]，如果不存在，就翻转整个数组；
// 2.再找出另一个最大索引 l 满足 nums[l] > nums[k]；
// 3.交换 nums[l] 和 nums[k]；
// 4.最后翻转 nums[k+1:]。
//
// 举个例子：
// 比如 nums = [1,2,7,4,3,1]，下一个排列是什么？
// 我们找到第一个最大索引是 nums[1] = 2
// 再找到第二个最大索引是 nums[4] = 3
// 交换，nums = [1,3,7,4,2,1];
// 翻转，nums = [1,3,1,2,4,7]
//
// 时间复杂度：O(n)
// 空间复杂度：O(1)

func nextPermutation(nums []int) {
	if nums == nil || len(nums) <= 1 {
		return
	}

	// 1.
	k := len(nums) - 2
	for ; k >= 0; k-- {
		if nums[k] < nums[k+1] {
			break
		}
	}

	// 若1.不存在，则反转整个数组 -- 此时说明数组已经是倒序排列了，最大
	if k < 0 {
		reverse(nums)
		return
	}

	// 2.
	l := len(nums) - 1
	for ; l >= 0; l-- {
		if nums[l] > nums[k] {
			break
		}
	}

	// 3.
	nums[k], nums[l] = nums[l], nums[k]

	// 4.
	reverse(nums[k+1:])
}

// 翻转数组
func reverse(nums []int) {
	if nums == nil || len(nums) <= 1 {
		return
	}

	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
