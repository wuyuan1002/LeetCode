package main

// 26. 删除有序数组中的重复项

// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次 ，返回删除后数组的新长度。
//
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

func main() {

}

// 双指针
func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	l, r := 0, 1
	for r < len(nums) {
		if nums[r] == nums[l] {
			r++
		} else {
			l++
			nums[l] = nums[r]
		}
	}

	return l + 1
}
