package main

// 80. 删除有序数组中的重复项 II

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素最多出现两次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

func main() {

}

// 类似 26、27
func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) < 3 {
		return len(nums)
	}

	l, r := 2, 2
	for r < len(nums) {
		if nums[r] != nums[l-2] {
			nums[l] = nums[r]
			l++
		}
		r++
	}
	return l
}
