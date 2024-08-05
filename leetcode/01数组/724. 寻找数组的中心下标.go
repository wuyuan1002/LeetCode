package main

// 724. 寻找数组的中心下标

// 给你一个整数数组 nums ，请计算数组的 中心下标 。
// 数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。
// 如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。
// 如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。

// pivotIndex .
// 先求出数组所有数字的总和（整个数组的右侧数字之和），然后从左到右遍历数组，计算每个下标的左右侧之和
func pivotIndex(nums []int) int {
	// 分别表示当前下标数字的左右侧之和
	leftSum, rightSum := 0, 0
	for _, n := range nums {
		rightSum += n
	}

	// 遍历数组，计算每个下标的左右侧之和，当一个数字的左右侧数字之和相同时说明找到了中心下标
	for i, n := range nums {
		rightSum -= n
		if leftSum == rightSum {
			return i
		}
		leftSum += n
	}

	return -1
}
