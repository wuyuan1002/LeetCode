package main

// 852. 山脉数组的峰顶索引
//
// 给定一个长度为 n 的整数 山脉 数组 arr ，其中的值递增到一个 峰值元素 然后递减。返回峰值元素的下标。
//
// 你必须设计并实现时间复杂度为 O(log(n)) 的解决方案。

// peakIndexInMountainArray .
// leetcode 153. 寻找旋转排序数组中的最小值
// leetcode 162. 寻找峰值
//
// 二分查找数组的峰顶 -- 数组在峰顶左侧是上升区域，在峰顶右侧是下降区域
// 这里需要使用 l < r 的二分查找，而不能使用 l <= r 的二分查找，
// 当nums[mid]比nums[mid+1]大时，说明mid到r是单调递减的，最大值一定在mid左侧（因为可能正好是mid，所以这里不能用 l <= r 的那种循环，也就不能使用 r=mid-1 这样赋值）
// 否则说明0到mid+1是单调递增的，最大值一定出现在mid右侧（因为0到mid是单调递增的且求的是最大值，此时mid一定不会是最大值，可能是mid+1，所以可以直接 l=mid+1 赋值）
func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1

	for l < r {
		mid := l + (r-l)/2
		// 若nums[mid] > nums[mid+1]，说明峰顶位于nums[0, mid]范围内（可能正好是nums[mid]），所以将r指向mid
		// 若nums[mid] <= nums[mid+1]，说明峰顶位于nums(mid, r]范围内（一定不会是nums[mid]），所以将l指向mid+1
		if arr[mid] > arr[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	// 二分查找结束后，l指向的就正好是最大值了
	return l
}
