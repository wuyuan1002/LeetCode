package main

import (
	"fmt"
)

// 239. 滑动窗口最大值

// 给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。
// 你只可以看到在滑动窗口内的 k个数字。滑动窗口每次只向右移动一位。
//
// 返回滑动窗口中的最大值。

func main() {
	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))
}

// 同 offer 59
// offer59那样做会超出时间限制
// 构建一个单调递减队列，具体实现类似于offer 59-2
// 见byte_dance 239的解法
func maxSlidingWindow(nums []int, k int) []int {

	if nums == nil || len(nums) == 0 || k <= 0 || k > len(nums) {
		return nil
	}

	maxqueue := make([]int, 0) // 单调递减队列 -- 存下标
	result := make([]int, 0)   // 结果集

	// 移动i形成窗口
	i := 0
	for ; i < k; i++ {
		index := 0
		if len(maxqueue) > 0 && nums[i] < nums[maxqueue[0]] {
			for ; index < len(maxqueue); index++ {
				if nums[maxqueue[index]] <= nums[i] {
					break
				}
			}
		}
		// 将nums[i]插入队列 -- nums[i]在队列末尾，是队列中的最小值
		maxqueue = append(maxqueue[:index], i)
	}

	// 形成窗口后，将第一个窗口的最大值入结果集
	result = append(result, nums[maxqueue[0]])

	// 形成窗口后，向后移动窗口
	for ; i < len(nums); i++ {

		// 删掉队列中超出窗口大小的元素 -- 若窗口中最大值超出窗口大小
		if len(maxqueue) > 0 && i-maxqueue[0]+1 > k {
			index := 0
			for ; index < len(maxqueue); index++ {
				if i-maxqueue[index]+1 <= k {
					break
				}
			}
			maxqueue = maxqueue[index:]
		}

		// 将新入窗口的元素加入到队列中
		index := 0
		if len(maxqueue) > 0 && nums[i] < nums[maxqueue[0]] {
			for ; index < len(maxqueue); index++ {
				if nums[maxqueue[index]] <= nums[i] {
					break
				}
			}
		}
		// 将nums[i]插入队列 -- nums[i]在队列末尾，是队列中的最小值
		maxqueue = append(maxqueue[:index], i)

		// 将队列中的最大值入结果集
		result = append(result, nums[maxqueue[0]])
	}

	return result
}
