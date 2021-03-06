package main

// 283. 移动零

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// func main() {

// }

// 1.使用两个队列，一个放非0元素，一个放0，依次遍历数组，将元素放入队列，之后将两个队列合并
// 2.双指针，一个指针始终指向第一个0，另一个指针向后遍历数组，找到非0元素，就和第一个指针交换
func moveZeroes(nums []int) {
	if nums == nil || len(nums) <= 1 {
		return
	}

	// 左指针左边均为非零数；
	// 右指针左边直到左指针处均为零。
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != 0 {
			// 若右面的指针指向的元素非0，交换两指针的元素
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}
