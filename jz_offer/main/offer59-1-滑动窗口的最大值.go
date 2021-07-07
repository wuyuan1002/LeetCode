package main

import (
	"fmt"
)

// 剑指 Offer 59 - I. 滑动窗口的最大值

// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

func main() {
	// fmt.Println(maxSlidingWindow1([]int{1, -1}, 1))
	fmt.Println(maxSlidingWindow1([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

// 同Hot100 239
func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil || len(nums) == 0 || k <= 0 {
		return nil
	}

	result := make([]int, 0)

	i := 0     // 窗口左下标 -- 始终指向最大值
	j := k - 1 // 窗口右下标

	for j < len(nums) {
		if i == 0 || j-i+1 > k { // 刚开始或者i一直是最大值，新进来值后超出了窗口大小 -- 重新计算最大值的下标
			if j-i >= k { // 若不是刚开始，则先把i向前移一位 -- 因为新进来了j向右移了一位，要保持窗口大小不变
				i++
			}
			for t := i; t <= j; t++ {
				if nums[t] > nums[i] {
					i = t
				}
			}
		} else if nums[j] >= nums[i] { // 新进来的值没有超出窗口大小, 且比i大或相等
			i = j
		}

		result = append(result, nums[i])
		j++
	}

	return result
}

// 第二次做
func maxSlidingWindow1(nums []int, k int) []int {
	if nums == nil || len(nums) < k || k <= 0 {
		return nil
	}

	result := make([]int, 0)
	i := 0
	maxqueue := []int{0}

	for ; i < len(nums); i++ {
		for j := 0; j < len(maxqueue); j++ {
			if i-maxqueue[j]+1 > k { // 判断是否超出窗口大小
				maxqueue = maxqueue[j+1:]
				j = 0
				continue
			}
			if nums[i] >= maxqueue[j] {
				maxqueue = maxqueue[:j]
				maxqueue = append(maxqueue, i)
				break
			}
			if j == len(maxqueue)-1 {
				maxqueue = append(maxqueue, i)
				break
			}
		}

		// if i-maxqueue[0]+1 > k {
		//
		// }

		result = append(result, maxqueue[0])
	}

	return result
}
