package 代码随想录

// 27. 移除元素

// 给你一个数组 nums和一个值 val，你需要 原地
// 移除所有数值等于val的元素，并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

// removeElement .
// 类似 26、80、283
// 双指针（快慢指针） -- 慢指针i之前的元素都是没有目标值的元素。i指向的是第一个等于val的值, i之前的是不等于val的所有值
func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	i, j := 0, 0 // 快慢指针
	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j] // 此处使用交换更加直观，只是等于val的值后续不再使用，所以可以直接赋值覆盖。nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	return i
}
