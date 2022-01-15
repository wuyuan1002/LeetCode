package main

// 215. 数组中的第K个最大元素

// 在未排序的数组中找到第 k 个最大的元素。
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

// func main() {
// 	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
// }

func findKthLargest(nums []int, k int) int {
	if nums == nil || len(nums) == 0 || k <= 0 || k > len(nums) {
		return 0
	}

	left, right := 0, len(nums)-1
	index := partition(nums, left, right)
	for index != k-1 {
		if index > k-1 {
			right = index - 1
		} else if index < k {
			left = index + 1
		}

		index = partition(nums, left, right)
	}

	return nums[index]
}

func partition(nums []int, left, right int) int {

	l, r := left, right
	tmp := nums[left]

	for l < r {
		for l < r && nums[r] <= tmp {
			r--
		}
		for l < r && nums[l] >= tmp {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

	nums[l], nums[left] = tmp, nums[l]

	return l
}

// --------------

// 第二次做
func findKthLargest1(nums []int, k int) int {
	if nums == nil || len(nums) == 0 || k <= 0 || k > len(nums) {
		return -1
	}

	l, r := 0, len(nums)-1
	index := partition1(nums, l, r)
	for index != k-1 {
		if index > k-1 {
			r = index - 1
		} else if index < k-1 {
			l = index + 1
		}

		index = partition1(nums, l, r)
	}

	return nums[index]
}

func partition1(nums []int, left, right int) int {
	l, r := left, right
	tmp := nums[left]
	for l < r {
		for l < r && nums[r] <= tmp {
			r--
		}
		for l < r && nums[l] >= tmp {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[left], nums[l] = nums[l], tmp

	return l
}
