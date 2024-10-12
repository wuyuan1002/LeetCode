package main

// 27. 移除元素

// 给你一个数组 nums和一个值 val，你需要 原地
// 移除所有数值等于val的元素，并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

// removeElement .
// leetcode 26. 删除有序数组中的重复项
// leetcode 80. 删除有序数组中的重复项 II
// leetcode 283. 移动零
//
// 双指针
// 左指针l左侧都是不等于val的值 -- [0, l)
// 左指针l指向的是第一个等于val的值 -- l
// 左指针l及其右侧至右指针r的左侧都是等于val的值 -- [l, r)
// 右指针r及其右侧都是正在遍历和还未遍历的元素 -- [r, len(nums)-1]
func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	// 左右指针
	l, r := 0, 0

	for r < len(nums) {
		// 若当前元素不等于val说明当前元素是目标值，要将其放到左指针左侧，所以将其赋值给左指针并将左指针向右移动一位
		if nums[r] != val {
			// 此处使用交换更加直观，只是等于val的值后续不再使用，所以可以直接赋值覆盖
			// nums[l], nums[r] = nums[r], nums[l]
			nums[l] = nums[r]
			l++
		}

		// 每次比较一个数字后都将右指针向右移动一位，开始遍历下一个数字
		r++
	}

	// 最终左指针左侧的数字都是不等于val的元素 -- [0, l)
	return l
}
