package main

// 350. 两个数组的交集 II

// 给定两个数组，编写一个函数来计算它们的交集。

// func main() {

// }

// 349
// 使用map存数字出现的次数
func intersect(nums1 []int, nums2 []int) []int {
	if nums1 == nil || len(nums1) == 0 || nums2 == nil || len(nums2) == 0 {
		return nil
	}

	hash := make(map[int]int, 0)
	for _, n := range nums1 {
		hash[n]++
	}

	result := make([]int, 0)
	for _, n := range nums2 {
		if _, ok := hash[n]; ok && hash[n] > 0 {
			result = append(result, n)
			hash[n]--
		}
	}

	return result
}

// 第二次做
func intersect1(nums1 []int, nums2 []int) []int {
	if nums1 == nil || len(nums1) == 0 || nums2 == nil || len(nums2) == 0 {
		return nil
	}

	hash := make(map[int]int)
	for _, n := range nums1 {
		hash[n]++
	}

	result := make([]int, 0)
	for _, n := range nums2 {
		if _, ok := hash[n]; ok && hash[n] > 0 {
			result = append(result, n)
			hash[n]--
		}
	}

	return result
}
