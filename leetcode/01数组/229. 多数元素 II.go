package main

// 229. 多数元素 II

// 给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

// majorityElement229 .
// leetcode 169. 多数元素
//
// 摩尔投票 -- 每三个不同的数字抵消一次，最终剩下的就是出现超过1/3的数字
func majorityElement229(nums []int) []int {
	element1, element2 := 0, 0 // 第一个元素和第二个元素
	vote1, vote2 := 0, 0       // 第一个元素投票数和第二个元素投票数

	for _, num := range nums {
		if vote1 > 0 && num == element1 { // 如果该元素为第一个元素，则计数加1
			vote1++
		} else if vote2 > 0 && num == element2 { // 如果该元素为第二个元素，则计数加1
			vote2++
		} else if vote1 == 0 { // 选择第一个元素
			element1 = num
			vote1++
		} else if vote2 == 0 { // 选择第二个元素
			element2 = num
			vote2++
		} else { // 如果三个元素均不相同，则相互抵消1次
			vote1--
			vote2--
		}
	}

	// 统计第一个元素和第二个元素出现的次数
	cnt1, cnt2 := 0, 0
	for _, num := range nums {
		if vote1 > 0 && num == element1 {
			cnt1++
		}
		if vote2 > 0 && num == element2 {
			cnt2++
		}
	}

	// 检测元素出现的次数是否满足要求
	result := make([]int, 0)
	if vote1 > 0 && cnt1 > len(nums)/3 {
		result = append(result, element1)
	}
	if vote2 > 0 && cnt2 > len(nums)/3 {
		result = append(result, element2)
	}

	return result
}
