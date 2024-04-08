package main

// 852. 山脉数组的峰顶索引

// 符合下列属性的数组 arr 称为 山脉数组 ：
// arr.length >= 3
// 存在 i（0 < i < arr.length - 1）使得：
// arr[0] < arr[1] < ... arr[i-1] < arr[i]
// arr[i] > arr[i+1] > ... > arr[arr.length - 1]
// 给你由整数组成的山脉数组 arr ，返回满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i 。
//
// 你必须设计并实现时间复杂度为 O(log(n)) 的解决方案。

// peakIndexInMountainArray .
// 同 leetcode 153. 寻找旋转排序数组中的最小值
// 二分查找数组的峰顶，数组在峰顶左侧是上升区域，在峰顶右侧是下降区域
func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1

	for l < r {
		mid := l + (r-l)/2
		// 当nums[mid]比nums[mid+1]大时，说明mid到r是单调递减的，最大值一定在mid左侧（因为可能正好是mid，所以这里不能用 l <= r 的那种循环，也就不能使用 r=mid-1 这样赋值）
		// 否则说明0到mid+1是单调递增的，最大值一定出现在mid右侧（因为0到mid是单调递增的且求的是最大值，此时mid一定不会是最大值，可能是mid+1，所以可以直接 l=mid+1 赋值）
		if arr[mid] > arr[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	// 二分查找结束后，l指向的就正好是最大值了
	return l
}
