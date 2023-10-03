package main

// 349. 两个数组的交集

// 给定两个数组 nums1 和 nums2 ，返回 它们的交集 。
// 输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。

// intersection .
// 同 350. 两个数组的交集 II
// 先将两个字符串分别存入set，进行字符去重并方便查找
func intersection(nums1 []int, nums2 []int) []int {

	// 两字符串存入set
	hash1 := make(map[int]struct{})
	for _, n := range nums1 {
		hash1[n] = struct{}{}
	}

	hash2 := make(map[int]struct{})
	for _, n := range nums2 {
		hash2[n] = struct{}{}
	}

	// 求两个set的交集
	result := make([]int, 0)
	for n := range hash2 {
		if _, ok := hash1[n]; ok {
			result = append(result, n)
		}
	}

	return result
}
