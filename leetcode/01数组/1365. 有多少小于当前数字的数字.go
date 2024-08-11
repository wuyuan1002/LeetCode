package main

// 1365. 有多少小于当前数字的数字

// 给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。
// 换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。
// 以数组形式返回答案。
//
// 提示：
// 2 <= nums.length <= 500
// 0 <= nums[i] <= 100

// smallerNumbersThanCurrent .
// 由于题目要求数组中的数字都在[0, 100]范围内，因此使用一个101个元素的数组用来记录每个数字的出现次数（下标为数字，对应的值为该数字出现的次数）
func smallerNumbersThanCurrent(nums []int) []int {
	// 构建一个数组用来存每个数字出现的次数
	count := make([]int, 101)

	// 遍历数组，统计每个数字出现的次数 -- 此时数组中每个数字的值为其下标数字出现的次数
	for _, n := range nums {
		count[n]++
	}

	// 遍历计数数组，从前向后累加小于等于每个下标数字的数字出现次数 -- 此时数组中每个数字的值为小于等于其下标数字出现的次数
	for i := 1; i < 101; i++ {
		count[i] += count[i-1]
	}

	// 统计结果，按照nums的数字顺序返回给定数字集合中小于每个数字的数字次数
	result := make([]int, 0, len(nums))
	for _, n := range nums {
		if n == 0 {
			// 小于0的数字出现的次数一定为0 -- 因为数字范围为[0, 100]
			result = append(result, 0)
		} else {
			// 其它小于n的数字出现的次数为count[n-1] -- count[n]存的是小于等于n的数字出现次数，需要将等于n的次数去掉 -- 即count[n-1]
			result = append(result, count[n-1])
		}
	}

	return result
}
