package main

import (
	"math"
)

// 16. 最接近的三数之和

// 给定一个包括n个整数的数组nums和 一个目标值target。
// 找出nums中的三个整数，使得它们的和与target最接近。
// 返回这三个数的和。假定每组输入只存在唯一答案。

func main() {

}

// 类似三数之和
func threeSumClosest(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	quickSort(nums, 0, len(nums)-1)

	result := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		// 若与上次确定的数相等，跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 先确定一个数
		num := nums[i]

		// 双指针遍历后面的数
		l, r := i+1, len(nums)-1
		for l < r {
			res := num + nums[l] + nums[r]
			if res == target {
				// 若与目标值相等了，直接返回
				return res
			}
			// 更新答案为绝对值差值最小的数
			if abs(res-target) < abs(result-target) {
				result = res
			}

			if res < target {
				l++
			} else {
				r--
			}
		}
	}

	return result
}

func quickSort(nums []int, left, right int) {
	if nums == nil || len(nums) == 0 || left >= right {
		return
	}
	l, r := left, right
	tmp := nums[left]
	for l < r {
		for l < r && nums[r] >= tmp {
			r--
		}
		for l < r && nums[l] <= tmp {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

	nums[l], nums[left] = tmp, nums[l]

	quickSort(nums, left, l-1)
	quickSort(nums, r+1, right)
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
