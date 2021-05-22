package main

// 给定一个整数数组 nums和一个整数目标值 target，
// 请你在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。

func main() {

}

// 可以在查找时将数字存入map中，这样，下次使用时可以直接获取
func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	hash := make(map[int]int) // k:数字，v:下标
	for i, num := range nums {

		// 若map中找到了target-num，说明存在两数之和等于target -- 返回下标
		if index, ok := hash[target-num]; ok {
			return []int{i, index}
		}

		// 若map中没有target-num，则将当前数字存入map，提供后面的数字查找
		hash[num] = i
	}
	return nil
}
