package main

// 33. 搜索旋转排序数组

func main() {

}

// offer 11
// 二分法
func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] { // 若左面是排好序的
			if nums[mid] > target && nums[0] <= target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // 若右面是排好序的
			if nums[mid] < target && nums[len(nums)-1] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}

// offer 11
// 二分法
func search1(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] {
			if nums[0] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[len(nums)-1] >= target && nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
