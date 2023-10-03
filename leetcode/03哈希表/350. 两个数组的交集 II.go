package main

// 350. 两个数组的交集 II
// 给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。
// 返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。
// 可以不考虑输出结果的顺序。

// intersect .
// 使用map存数字出现的次数
func intersect(nums1 []int, nums2 []int) []int {
	// 遍历nums1 -- 将每个值及其出现的次数存入map
	hash1 := make(map[int]int)
	for _, n := range nums1 {
		hash1[n]++
	}

	// 遍历nums2 -- 判断hash1中数字是否还有余量
	result := make([]int, 0)
	for _, n := range nums2 {
		if count, ok := hash1[n]; ok && count > 0 {
			result = append(result, n)
			hash1[n]--
		}
	}

	return result
}
