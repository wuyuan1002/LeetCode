package main

// 4. 寻找两高正序数组的中位数

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
// 请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。

// findMedianSortedArrays .
// 1. 合并两个数组，求中位数 -- 合并过程类似归并排序的一部分，见Offer 51
// 2. 不需要合并两个数组，由于两个数组长度已知，因此最终的中位数是第几个也是已知的，使用两个指针同时遍历两个数组即可得到中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	midNum := 0                       // 中位数
	alllen := len(nums1) + len(nums2) // 总数字个数
	midIndex := alllen / 2            // 中位数是总数字中的第几个
	next := 0                         // 中位数的下一个数字 -- 若总数为偶数时使用

	// 若总长度为奇数，前面求的 alllen/2 会默认向下取整，实际奇数时的中位数需要+1
	if alllen%2 != 0 {
		midIndex += 1
	}

	// 双指针遍历两个数组，两指针向后移动, 若有一个已到达末尾, 则指移动另一个
	// 因为中位数位于总数的第midIndex个，而每次遍历都将双指针向后移动一位，那么这个循环循环第midIndex次时双指针所指向的数字就是中位数，所以这里使用midIndex限制循环次数
	// 这里允许midIndex等于0是因为在找到中位数后若总数字个数为偶数个则需要继续寻找中位数的下一个数字，所以允许这个循环执行midIndex+1次，多循环一次来找到中位数的下一个数字
	for i, j := 0, 0; midIndex >= 0; midIndex-- {
		// 若已经找到中位数，且总数字个数为奇数，说明没必要寻找中位数的下一个数字，直接break返回已经找到的中位数即可
		if midIndex == 0 && alllen%2 != 0 {
			break
		}

		// 双指针向后移动, 若有一个已到达末尾, 则只移动另一个
		if i == len(nums1) { // i已指向数组末尾，只向后移动j
			if midIndex == 1 { // 找到中位数
				midNum = nums2[j]
			} else if midIndex == 0 { // 找中位数的下一个数字
				next = nums2[j]
			}
			j++
		} else if j == len(nums2) { // j已指向数组末尾，只向后移动i
			if midIndex == 1 { // 找到中位数
				midNum = nums1[i]
			} else if midIndex == 0 { // 找中位数的下一个数字
				next = nums1[i]
			}
			i++
		} else { // 若两个指针都未到达末尾，比较双指针指向数字的大小，向后移动数字小的那个指针
			if nums1[i] <= nums2[j] {
				if midIndex == 1 { // 找到中位数
					midNum = nums1[i]
				} else if midIndex == 0 {
					next = nums1[i]
				}
				i++
			} else {
				if midIndex == 1 { // 找到中位数
					midNum = nums2[j]
				} else if midIndex == 0 { // 找中位数的下一个数字
					next = nums2[j]
				}
				j++
			}
		}
	}

	// 若总个数为偶数，返回中位数与下一个数字的平均数，若总数是奇数，直接返回中位数
	if alllen%2 == 0 {
		return float64(midNum+next) / 2
	} else {
		return float64(midNum)
	}

}
