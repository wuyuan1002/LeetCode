package main

import "sort"

// 435. 无重叠区间

// 给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。
// 返回 需要移除区间的最小数量，使剩余区间互不重叠 。

// eraseOverlapIntervals .
// leetcode 452. 用最少数量的箭引爆气球
//
// 移除最少区间使得剩余的不重叠，相当于求有多少互不重叠的区间，再用总数减去即可
// 将所有区间按照右端点从小到大排序，然后遍历排好序的区间右端点进行区间选择，选择出其中不重叠的区间
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}

	// 按区间右端点从小到大排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })

	// 最多的不重复区间个数、当前遍历过程中的最右侧端点下标
	notOverlap, currentRight := 1, intervals[0][1]

	// 不断更新当前的最右侧端点下标，选择左端点大于当前右端点的区间的右端点作为新的当前最右侧端点下标
	for _, interval := range intervals[1:] {
		// 左端点与右端点相同时两区间算为不重叠
		// 若当前区间的左侧端点大于或等于当前右侧端点下标，说明当前区间是最近的与前一个区间不重叠的区间
		// 递增不重叠区间的个数并更新不重叠区间的最右侧端点下标为当前区间的右端点
		if interval[0] >= currentRight {
			notOverlap++
			currentRight = interval[1]
		}
	}

	// 总数减去最多不重叠区间的个数就是需要移除的区间个数
	return len(intervals) - notOverlap
}
