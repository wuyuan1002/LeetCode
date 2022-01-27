package main

// 349. 两个数组的交集

// 给定两个数组，编写一个函数来计算它们的交集。

// func main() {
// 	fmt.Println(intersection1([]int{1}, []int{1}))
// }

// 350
// 使用map存数字是否出现过
func intersection(nums1 []int, nums2 []int) []int {
	if nums1 == nil || len(nums1) == 0 || nums2 == nil || len(nums2) == 0 {
		return nil
	}

	hash1 := make(map[int]struct{}, 0)
	for _, n := range nums1 {
		hash1[n] = struct{}{}
	}

	hash2 := make(map[int]struct{}, 0)
	for _, n := range nums2 {
		hash2[n] = struct{}{}
	}

	result := make([]int, 0)
	for n := range hash2 {
		if _, ok := hash1[n]; ok {
			result = append(result, n)
		}
	}

	return result
}

// 第二次做
func intersection1(nums1 []int, nums2 []int) []int {
	if nums1 == nil || len(nums1) == 0 || nums2 == nil || len(nums2) == 0 {
		return nil
	}

	hash1 := make(map[int]bool)
	for _, n := range nums1 {
		hash1[n] = true
	}

	hash2 := make(map[int]bool)
	for _, n := range nums2 {
		hash2[n] = true
	}

	result := make([]int, 0)
	for n := range hash2 {
		if _, ok := hash1[n]; ok {
			result = append(result, n)
		}
	}

	return result
}
