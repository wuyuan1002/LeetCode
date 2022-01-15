package main

// 75. 颜色分类

// 给定一个包含红色、白色和蓝色，一共n 个元素的数组，
// 原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
// 此题中，我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。

// func main() {
// 	aa := []int{0, 1, 0, 0, 1}
// 	sortColors(aa)
// 	fmt.Println(aa)
// }

// 插入排序
// 双指针 p0、p2
func sortColors(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}

	// p0左面的都是0
	// p2右面的都是2
	// p0、i之间的都是1

	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			i++
		} else if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		} else {
			panic("数组中元素只能为0、1、2")
		}
	}
}
