package main

// 169. 多数元素

// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

// majorityElement .
// leetcode 229. 多数元素 II
//
// 摩尔投票、求水王数
//
// 相当于一次删除两个不同的数字，有水王数时最后剩下的就是水王数
// 但若没有水王数时，也可能剩下来，此时需要再遍历一次数组查看剩下来数字的出现次数，比如 [1, 2, 3, 4, 5] 会剩下 5
func majorityElement(nums []int) int {
	num, veto := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if veto == 0 {
			num = nums[i]
		}

		if nums[i] == num {
			veto++
		} else {
			veto--
		}
	}

	return num
}
