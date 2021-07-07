package main

// 912. 排序数组

// 给你一个整数数组 nums，请你将该数组升序排列。

func main() {

}

func sortArray(nums []int) []int {
	quickSort2(nums, 0, len(nums)-1)
	return nums
}

func quickSort2(arr []int, left, right int) {
	if arr == nil || len(arr) == 0 || left >= right {
		return
	}

	l, r := left, right
	tmp := arr[left]
	for l < r {
		for l < r && arr[r] >= tmp {
			r--
		}
		for l < r && arr[l] <= tmp {
			l++
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	arr[l], arr[left] = tmp, arr[l]
	quickSort2(arr, left, l-1)
	quickSort2(arr, r+1, right)
}
