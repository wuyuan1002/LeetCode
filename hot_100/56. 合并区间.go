package main

import (
	"fmt"
	"sort"
)

// 56. 合并区间

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi].
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}

func merge(intervals [][]int) [][]int {
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

	// 总结果
	result := make([][]int, 0)
	// 当前总区间
	current := intervals[0]

	// 遍历每一个区间
	var interval []int
	for i := 1; i < len(intervals); i++ {
		interval = intervals[i]
		// 该区间与当前总区间有交集
		if interval[0] <= current[1] {
			// 更新当前总区间的最大右边界
			current[1] = max(interval[1], current[1])
		} else {
			// 新的区间与当前总区间没有交集
			// 将当前总区间写入总结果
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)

			// 把当前总区间设为当前区间
			current = interval
		}
	}
	// 写入最后一个总区间结果
	result = append(result, current)

	return result
}
