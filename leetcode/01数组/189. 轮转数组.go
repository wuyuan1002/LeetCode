package main

// 189. 轮转数组

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

// rotate .
// 同 leetcode 151. 反转字符串中的单词
func rotate(nums []int, k int) {
	// 计算真正需要旋转的下标，若为0说明不需要旋转
	k = k % len(nums)
	if k == 0 {
		return
	}

	// 先反转整个数组，然后以下标k为分界反转[0, k)和[k, len(nums)-1]的数字
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
