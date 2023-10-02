package main

// 162. 寻找峰值

// func main() {

// }

// 1. 一次遍历数组，寻找比它前一个和后一个都大的元素下标 O(n)
// 2. 二分查找 O(log n)
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

// 二分查找 -- 只要数组中存在一个元素比相邻元素大，那么沿着它一定可以找到一个峰值
func findPeakElement1(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2 // 防止l+r时超出int的范围
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
