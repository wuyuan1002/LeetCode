package main

import "sort"

// 57. 插入区间

// 给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
// 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

// insert .
// 遍历所有区间，若区间和插入区间没有交集则直接加入结果集，若与插入区间有交集则将其与插入区间进行合并，
// 最后将合并后的插入区间加入到结果集，然后按照左端点进行排序后返回
func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)                    // 总结果
	left, right := newInterval[0], newInterval[1] // 插入区间的左右范围

	// 遍历所有区间，判断其与要插入的区间是否有重叠，若有重叠则将其与要插入的区间进行合并，若没有重叠则将该区间直接加入到结果集
	for _, interval := range intervals {
		if interval[1] < left || interval[0] > right {
			// 若该区间与插入区间无交集 -- 直接将区间入结果集
			result = append(result, interval)
		} else {
			// 将当前区间和插入区间进行合并，求出新区间的左右范围
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}

	// 将合并后的插入区间加入到结果集
	result = append(result, []int{left, right})

	// 将所有区间按照左端点排序 -- 因为最后插入的插入区间是在结果的最后一个，题目要求按照区间顺序返回
	sort.Slice(result, func(i, j int) bool { return result[i][0] < result[j][0] })

	return result
}
