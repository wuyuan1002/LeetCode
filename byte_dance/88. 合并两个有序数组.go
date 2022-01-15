package main

import (
	"math"
)

// 88. 合并两个有序数组

// 给你两个有序整数数组nums1 和 nums2，请你将 nums2 合并到nums1中，使 nums1 成为一个有序数组。
// 初始化nums1 和 nums2 的元素数量分别为m 和 n 。
// 你可以假设nums1 的空间大小等于m + n，这样它就有足够的空间保存来自 nums2 的元素。

// func main() {

// }

// 从后向前同时遍历两个数组，将两数组末尾大拿数字放到第一个数组的最后面
func merge(nums1 []int, m int, nums2 []int, n int) {
	if nums1 == nil || nums2 == nil {
		return
	}

	end := m + n - 1 // num1的最后一个位置
	i, j := m-1, n-1 // i,j先分别指向两数组的最后一个数字
	ni, nj := 0, 0   // 两数组在i,j位置上的数字
	for i >= 0 || j >= 0 {
		// 获取要比较的数字
		if i >= 0 {
			ni = nums1[i]
		} else {
			ni = math.MinInt32
		}
		if j >= 0 {
			nj = nums2[j]
		} else {
			nj = math.MinInt32
		}

		// 将大的数字放到数组末尾
		if ni > nj {
			nums1[end] = ni
			i--
		} else {
			nums1[end] = nj
			j--
		}

		// 将末尾指针向前移动一位
		end--
	}
}
