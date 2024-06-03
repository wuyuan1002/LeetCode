package main

// 15. 三数之和

// 给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，
// 使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。

// threeSum .
// 先固定一个值，接下来就是在剩余数字中寻找两数之和，但是需要保证结果不重复，
// 如果寻找两数之和时使用map的话，结果会重复；因此我们先将数组排序，在寻找两数之和时使用双指针从两边移动来找
// 排序 + 双指针
func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}

	quickSort(nums, 0, len(nums)-1)

	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			// 剪枝 -- 若i指向大于0的数，则直接退出，因为数组已排序，i后面的数也肯定都是正数
			break
		} else if i > 0 && nums[i] == nums[i-1] {
			// 剪枝 -- 跳过重复元素
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				// 记录结果
				result = append(result, []int{nums[i], nums[l], nums[r]})

				// 剪枝 -- 跳过重复元素
				for l++; l < r && nums[l] == nums[l-1]; l++ {
				}
				for r--; l < r && nums[r] == nums[r+1]; r-- {
				}
			}
		}
	}

	return result
}
