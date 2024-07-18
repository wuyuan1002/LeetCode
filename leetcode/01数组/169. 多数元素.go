package main

// 169. 多数元素

// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

// majorityElement .
// leetcode 229. 多数元素 II
//
// 摩尔投票 -- 每两个不同的数字抵消一次，最终剩下的就是出现超过1/2的数字
func majorityElement(nums []int) int {
	// 目标数字、目标数字投票数
	element, veto := 0, 0

	for _, num := range nums {
		// 若目标数字的投票数为0则选择当前数字为目标数字
		if veto == 0 {
			element = num
		}

		// 若当前数字为目标数字则增加目标数字的投票数，否则将当前数字和目标数字进行抵消（递减目标数字的投票数）
		if num == element {
			veto++
		} else {
			veto--
		}
	}

	return element
}
