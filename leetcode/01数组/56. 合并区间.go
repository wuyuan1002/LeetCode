package main

import "sort"

// 56. 合并区间

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi].
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

// merge .
// 将每个区间按照左端点升序排序，然后进行区间合并，
// 若后一个区间的左端点大于前一个区间的右端点，说明两个区间没有交集，否则说明有交集
func merge(intervals [][]int) [][]int {
	// 按照左端点升序排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	result := make([][]int, 0) // 总结果
	current := intervals[0]    // 当前合并区间

	// 从第2个区间开始进行区间合并
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i] // 当前区间

		if interval[0] <= current[1] {
			// 当前区间与当前合并区间有交集 -- 更新当前合并区间的最大右边界
			current[1] = max(interval[1], current[1])
		} else {
			// 当前区间与当前合并区间没有交集 -- 将当前合并的总区间记录到结果集
			result = append(result, []int{current[0], current[1]})

			// 将当前区间置为下一个合并总区间，继续合并后续区间
			current = interval
		}
	}

	// 将最后一个合并区间加入到结果集
	result = append(result, current)

	return result
}
