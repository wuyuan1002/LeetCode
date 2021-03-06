package main

// 15. 三数之和

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
// 使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。

// func main() {

// }

// 1. 先排序数组，然后先固定一个值，然后在之后元素中使用双指针寻找两数之和
func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	res := make([][]int, 0)

	quickSort(nums, 0, len(nums)-1)

	for i, n := range nums {
		if n > 0 {
			// 剪枝
			break
		} else if i > 0 && n == nums[i-1] {
			// 跳过重复元素
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := n + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{n, nums[l], nums[r]})

				for l++; l < r && nums[l] == nums[l-1]; l++ {
				}
				for r--; l < r && nums[r] == nums[r+1]; r-- {
				}
			}
		}
	}

	return res
}

func threeSum1(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	quickSort(nums, 0, len(nums)-1)
	res := make([][]int, 0)
	for i, num := range nums {
		if num > 0 {
			break
		} else if i > 0 && num == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := num + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{num, nums[l], nums[r]})
				for l++; l < r && nums[l] == nums[l-1]; l++ {
				}
				for r--; l < r && nums[r] == nums[r+1]; r-- {
				}
			}
		}
	}

	return res
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
