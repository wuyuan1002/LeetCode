package main

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 进阶：你可以设计并实现时间复杂度为O(n) 的解决方案吗？

func main() {

}

// 1.使用插入排序，在排序过程中计算最长连续序列
// 2.依次遍历数组中每一个数字x，不断尝试寻找x+1,x+2...是否存在，这个寻找的过程，可以使用map来做到O(1)的时间复杂度
func longestConsecutive(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	// 将数组元素写入map中，同时进行了数字去重
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}

	longest := 0 // 最长连续长度
	for num := range numMap {
		if numMap[num-1] {
			// 开始查找的数字使用所在序列中的最小值，不然会做无用计算，
			// 如连续序列为x,x+1,x+2,x+3,x+4，我们却计算了x+2开始的序列长度，
			// 这个长度一定不会时最长序列长度，因为我们还会从x开始计算序列长度

			// 所以，如果存在比当前元素小1的元素，直接跳过当前元素，应该从更小的元素开始计算
			// 这里必须判断是小1的元素，如果时小2，那么就构不成连续序列了
			continue
		}

		// 依次遍历map，查找x+1,x+2,x+3,x+4...
		i := 1
		for ; ; i++ {
			if !numMap[num+i] {
				break
			}
		}

		// 更新最大序列长度
		longest = max(longest, i)
	}

	return longest
}
