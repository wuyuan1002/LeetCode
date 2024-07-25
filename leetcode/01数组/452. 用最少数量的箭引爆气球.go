package main

import "sort"

// 452. 用最少数量的箭引爆气球

// 有一些球形气球贴在一堵用 XY 平面表示的墙面上。墙面上的气球记录在整数数组 points ，其中points[i] = [xstart, xend] 表示水平直径在 xstart 和 xend之间的气球。你不知道气球的确切 y 坐标。
//
// 一支弓箭可以沿着 x 轴从不同点 完全垂直 地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被 引爆 。可以射出的弓箭的数量 没有限制 。 弓箭一旦被射出之后，可以无限地前进。
//
// 给你一个数组 points ，返回引爆所有气球所必须射出的 最小 弓箭数 。

// findMinArrowShots .
// leetcode 435. 无重叠区间
//
// 等于是求所有区间中有多少个互不重叠的区间
func findMinArrowShots(points [][]int) int {
	if len(points) <= 1 {
		return 1
	}

	// 按区间右端点从小到大排序
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })

	// 不重复区间个数、当前遍历过程中的最右侧端点下标
	// 不断更新当前的最右侧端点下标，选择左端点大于当前右端点的区间的右端点作为新的当前最右侧端点下标
	notOverlap, currentRight := 1, points[0][1]
	for _, interval := range points[1:] {
		// 左端点与右端点相同时两区间算为重叠
		// 若当前区间的左侧端点大于当前右侧端点下标，说明当前区间是最近的与前一个区间不重叠的区间
		// 递增不重叠区间的个数并更新不重叠区间的最右侧端点下标为当前区间的右端点
		if interval[0] > currentRight {
			notOverlap++
			currentRight = interval[1]
		}
	}

	return notOverlap
}
