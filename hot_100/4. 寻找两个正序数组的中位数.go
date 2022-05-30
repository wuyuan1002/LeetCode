package main

import "fmt"

// 4.寻找两高正序数组的中位数

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
// 请你找出并返回这两个正序数组的中位数 。

// func main() {
// 	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
// }

// 1. 合并两个数组，求中位数 -- 合并过程类似归并排序的一部分，见offer 51
// 2. 不需要合并两个数组，由于两个数组长度已知，因此最终的中位数是第几个也是已知的，使用两个指针同时遍历两个数组即可得到中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if nums1 == nil && nums2 == nil {
		return 0.0
	}
	if nums1 == nil {
		nums1 = make([]int, 0)
	}
	if nums2 == nil {
		nums2 = make([]int, 0)
	}

	alllen := len(nums1) + len(nums2) // 总数字个数
	isEven := alllen%2 == 0           // 总数是否为偶数
	middle := alllen / 2              // 中位数是总数字中的第几个
	mid := 0                          // 中位数
	next := 0                         // 中位数的下一个数字 -- 若总数为偶数时使用

	if !isEven {
		middle++
	}

	i, j := 0, 0      // 两个指针分别遍历两个数组
	for middle >= 0 { // 这里允许等于0是因为在找到中位数后继续寻找中位数的下一个数字
		if j == len(nums2) { // 若j已经到达数组末尾 -- 则只移动i
			if middle == 1 { // 已找到中位数
				fmt.Printf("00 %d", mid)
				mid = nums1[i]
			} else if middle == 0 { // 找中位数的下一个数字
				next = nums1[i]
			}
			i++
		} else if i == len(nums1) { // 若i已经到达数组末尾 -- 则只移动j
			if middle == 1 {
				fmt.Printf("11 %d", mid)
				mid = nums2[j]
			} else if middle == 0 {
				next = nums2[j]
			}
			j++
		} else { // 若两个指针都未到达数组末尾
			if nums1[i] <= nums2[j] {
				if middle == 1 {
					mid = nums1[i]
					fmt.Printf("22  %d", mid)
				} else if middle == 0 {
					next = nums1[i]
				}
				i++
			} else {
				if middle == 1 {
					fmt.Printf("%d -- %d", i, j)
					fmt.Printf("33 %d", mid)
					mid = nums2[j]
				} else if middle == 0 {
					next = nums2[j]
				}
				j++
			}
		}

		middle--

		// 若总数为奇数且已经找到中位数，则直接返回，不用再继续寻找下一个数字
		if !isEven && middle == 0 {
			break
		}
	}

	if isEven {
		return float64(mid+next) / 2
	} else {
		return float64(mid)
	}
}

func main() {
	fmt.Printf("%f", findMedianSortedArrays([]int{1, 3}, []int{2}))
}
