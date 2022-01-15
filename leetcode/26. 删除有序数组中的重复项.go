package main

// 26. 删除有序数组中的重复项

// func main() {

// }

// 双指针
func removeDuplicates1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	l, r := 0, 0
	for r < len(nums) {
		if nums[r] != nums[l] {
			l++
			nums[l] = nums[r]
		}
		r++
	}
	return l + 1
}
