package main

// 581. 最短无序连续子数组

// 给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，
// 那么整个数组都会变为升序排序。
//
// 请你找出符合题意的 最短子数组，并输出它的长度。

// func main() {
// 	fmt.Println(findUnsortedSubarray([]int{1, 3, 5, 2, 4}))
// }

// 先将数组排序，之后使用双指针比较两个数组不同的位置
func findUnsortedSubarray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	// 构建排序数组并进行排序
	sortedNum := make([]int, len(nums))
	copy(sortedNum, nums)
	quickSort1(sortedNum, 0, len(sortedNum)-1)

	// 遍历两数组，寻找无序子数组
	l, r := len(nums), 0 // 无序数组的起始和结束下标
	for i := 0; i < len(nums); i++ {
		if sortedNum[i] != nums[i] {
			l = min(l, i)
			r = max(r, i)
		}
	}

	if l >= r {
		return 0
	}

	return r - l + 1
}

func quickSort1(nums []int, left, right int) {
	if left >= right || nums == nil || len(nums) == 0 {
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

	quickSort1(nums, left, l-1)
	quickSort1(nums, r+1, right)
}
