package main

// 1. 两数之和

// 给定一个整数数组 nums和一个整数目标值 target，
// 请你在该数组中找出和为目标值 target 的那两个整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。

// twoSum .
// 类似 leetcode 202. 快乐数
// 在查找的过程中将数字存入map中，这样，下次使用时可以直接获取
func twoSum(nums []int, target int) []int {
	// 使用一个map存已经出现过的数字 -- k: 数值, v: 下标
	hash := make(map[int]int)

	for i, num := range nums {
		// 若target-num已经在map中，说明存在两个数字之和为target，直接将其返回即可
		if index, ok := hash[target-num]; ok {
			return []int{i, index}
		}
		// 将当前数字存入map供后续使用
		hash[num] = i
	}

	return nil
}
