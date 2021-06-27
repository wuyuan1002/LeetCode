package main

// 1. 两数之和

// 给定一个整数数组 nums和一个整数目标值 target，
// 请你在该数组中找出和为目标值 target 的那两个整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。

func main() {

}

func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	hash := make(map[int]int) // k: 数值, v: 下标
	for i, num := range nums {
		if index, ok := hash[target-num]; ok {
			return []int{i, index}
		}
		hash[num] = i
	}

	return nil
}
