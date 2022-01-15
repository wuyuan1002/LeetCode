package main

// 442. 数组中重复的数据

// 给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
//
// 找到所有出现两次的元素。
//
// 你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？

// func main() {
// 	fmt.Println(findDuplicates1([]int{4, 3, 2, 7, 8, 2, 3, 1}))
// }

// 1. 遍历数组，存入map的同时检查出现次数
// 2. offer 03
func findDuplicates(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	result := make([]int, 0)
	nowNum := 0 // 当前数字
	for i := 0; i < len(nums); {
		nowNum = nums[i]
		if nowNum > len(nums) {
			// 若当前数字已被访问过
			i++
			continue
		}

		if nowNum-1 == i {
			// 当前数字在正确的位置上
			i++
			continue
		}

		if nowNum == nums[nowNum-1] {
			result = append(result, nowNum)

			// 加上数组长度，标记该数已被记录
			nums[i] += len(nums)
			nums[nowNum-1] += len(nums)

			i++
			continue
		}

		nums[nowNum-1], nums[i] = nums[i], nums[nowNum-1]
	}

	// 将数组元素复原
	for i, n := range nums {
		if n > len(nums) {
			nums[i] -= len(nums)
		}
	}

	return result
}

func findDuplicates1(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	res := make([]int, 0) // 总结果
	i := 0                // 下标
	n := -1               // 当前数字
	for i < len(nums) {
		n = nums[i]
		if n > len(nums) {
			// 若已被访问过
			i++
			continue
		}
		if n-1 != i { // 若当前数字不在正确位置
			if n != nums[n-1] {
				nums[i], nums[n-1] = nums[n-1], nums[i]
				continue
			}

			// 记录重复数字并标记已被访问
			res = append(res, n)
			nums[i] += len(nums)
			nums[n-1] += len(nums)
		}
		i++
	}

	// 复原标记的元素
	for i := range nums {
		if nums[i] > len(nums) {
			nums[i] -= len(nums)
		}
	}

	return res
}
