package main

// 229. 求众数 II

// 给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

// func main() {
// 	fmt.Println(majorityElement([]int{2, 1, 1, 3, 1, 4, 5, 6}))
// }

// 超过n/k次的数字个数最多为k-1个

// 1. 使用map统计数字出现的次数
// 2. 摩尔投票
// 求水王数
// 每次删除3个不同的数，最后剩下的数就有可能是出现次数超过 n/3 次的数字
func majorityElement(nums []int) []int {
	res := make([]int, 0)
	if nums == nil || len(nums) == 0 {
		return res
	}

	// 由于求超过 n/3 次的数字，所以数字个数最多为2个
	num1, num2 := 0, 0   // 初始化两个候选者
	vote1, vote2 := 0, 0 // 两个候选者票数

	for _, num := range nums {
		if vote1 > 0 && num == num1 { // 若当前数是第一个数
			vote1++
		} else if vote2 > 0 && num == num2 { // 若当前数是第二个数
			vote2++
		} else if vote1 == 0 {
			num1 = num
			vote1++
		} else if vote2 == 0 {
			num2 = num
			vote2++
		} else {
			vote1--
			vote2--
		}
	}

	// 统计num1和num2在数组中实际出现的次数，判断是否真的众数
	count1, count2 := 0, 0
	for _, n := range nums {
		if n == num1 {
			count1++
		} else if n == num2 {
			count2++
		}
	}

	// 返回结果集
	if count1 > len(nums)/3 {
		res = append(res, num1)
	}
	if count2 > len(nums)/3 {
		res = append(res, num2)
	}

	return res
}
