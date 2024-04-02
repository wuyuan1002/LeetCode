package main

import "math"

// 581. 最短无序连续子数组

// 给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//
// 请你找出符合题意的 最短 子数组，并输出它的长度。

// findUnsortedSubarray .
// 数组可分为左中右3段，其中左和右段是排好序的，中段是乱序的，题目要求的就是中段的长度，
// 因为整个数组要是有序的，所以整体排序后，左段和右段的相对位置是不变的，所以一定满足中段的最小值大于左段的最大值，且中段的最大值小于右段的最小值。
//
// 从左到右看，数应该越来越大，如果某个数，比前面最大值大，那他没问题，如果小于前面的最大值，那么这个数就有问题，
// 从左到右不断更新最大值，遇到有问题的数，就记录下来作为right，并且right可以更新，
// right更新到最后面不动了，说明right右边的数都比，right左边的max的数大，但是right自己比max小，right右边是排好序的，right左边需要重新排序，right是需要重新排序的区间的右边界，
// left同理，left就是从右往左看，数要越来越小才行，如果某个数，比右边最小值还要小，那他没问题，如果它比右边的最小值要大，说明他有问题。
func findUnsortedSubarray(nums []int) int {
	// 遍历过的数字中出现的最小值和最大值、中段的左右边界
	min, max := math.MaxInt64, math.MinInt64
	left, right := -1, -1

	// 从左到右遍历数组 -- 确定中段右边界
	for i := 0; i <= len(nums)-1; i++ {
		// 更新右边界或最大值
		if nums[i] < max {
			right = i
		} else {
			max = nums[i]
		}
	}

	// 从右到左遍历数组 -- 确定中段左边界
	for i := len(nums) - 1; i >= 0; i-- {
		// 更新左边界或最小值
		if nums[i] > min {
			left = i
		} else {
			min = nums[i]
		}
	}

	// 返回中段子数组长度
	if left == -1 || right == -1 {
		return 0
	}
	return right - left + 1
}
