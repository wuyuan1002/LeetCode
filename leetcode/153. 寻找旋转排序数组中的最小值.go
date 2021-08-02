package main

// 153. 寻找旋转排序数组中的最小值

func main() {

}

// 二分法
// 同offer11，但是此题说了元素互不相同，因此比offer11要简单很多
func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return nums[l]
}
