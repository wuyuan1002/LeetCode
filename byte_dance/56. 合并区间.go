package main

import (
	"sort"
)

// 56. 合并区间

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi]。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

func main() {

}

// 类似57
func merge1(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 0 {
		return nil
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	// 按照左端点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0) // 总结果
	current := intervals[0]    // 当前合并总区间
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i] // 当前要合并的区间
		if interval[0] <= current[1] {
			// 更新总区间的右边界
			current[1] = max(current[1], interval[1])
		} else {
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			current = interval
		}
	}

	result = append(result, current)
	return result
}
