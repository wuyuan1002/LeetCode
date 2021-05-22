package main

// 剑指 Offer 57. 和为s的两个数字 -- 和leetcode第一题相似

// 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。
// 如果有多对数字的和等于s，则输出任意一对即可。

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}

func twoSum(nums []int, target int) []int {

	result := make([]int, 2)
	if nums == nil || len(nums) == 0 {
		return result
	}

	left := 0              // 左指针
	right := len(nums) - 1 // 右指针
	sum := 0               // 两数之和
	for left < right && left >= 0 && right < len(nums) {
		sum = nums[left] + nums[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			result[0] = nums[left]
			result[1] = nums[right]
			break
		}
	}
	return result
}

// 第二次做 -- 由于数组已排序，所以可以使用两个指针遍历数组
func twoSum1(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	l := 0
	r := len(nums) - 1
	sum := 0
	for l < r {
		sum = nums[l] + nums[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			return []int{nums[l], nums[r]}
		}
	}
	return nil
}

// LeetCode 1: 两数之和 -- 只是这个数组没有排好序
// 可以在查找时将数字存入map中，这样，下次使用时可以直接获取
func twoSum3(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	hash := make(map[int]int) // k:数字，v: 下标
	for i, num := range nums {
		if index, ok := hash[target-num]; ok { // 若map中存在要查找的数字，直接返回
			return []int{index, i}
		}
		hash[num] = i // 若map中不存在要找的数字，将当前数字存入map，以便别的数字查找
	}
	return nil
}
