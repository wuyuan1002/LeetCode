package main

import (
	"fmt"
)

// 给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，
// 使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。

func main() {
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
}

// 先固定一个值，接下来就是在剩余数字中寻找两数之和，但是需要保证结果不重复，
// 如果寻找两数之和时使用map的话，结果会重复；因此我们先将数组排序，在寻找两数之和时使用双指针从两边移动来找
// 排序 + 双指针
func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}

	result := make([][]int, 0)

	// 排序
	quickSort(nums, 0, len(nums)-1)

	target := 0
	for i := 0; i < len(nums)-1; i++ {

		// 若i指向大于0的数，则直接退出，因为数组已排序，i后面的数也肯定都是正数
		if nums[i] > 0 {
			break
		}

		// 跳过重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target = -nums[i]
		l, r := i+1, len(nums)-1 // 左右指针
		for l < r {
			if nums[l]+nums[r] < target {
				l++
			} else if nums[l]+nums[r] > target {
				r--
			} else {
				// 只将与上一次结果不相等的值存入结果中 -- 见 -2, 0, 0, 2, 2
				if l > i+1 && r < len(nums)-1 && nums[l] == nums[l-1] && nums[r] == nums[r+1] {
					l++
					r--
					continue
				}
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			}
		}
	}

	return result
}

// 快排
func quickSort(nums []int, left, right int) {
	if left >= right || nums == nil || len(nums) == 0 {
		return
	}

	l, r := left, right
	temp := nums[left]
	for l < r {
		for l < r && nums[r] >= temp {
			r--
		}
		for l < r && nums[l] <= temp {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[left], nums[l] = nums[l], temp

	quickSort(nums, left, l-1)
	quickSort(nums, r+1, right)
}
