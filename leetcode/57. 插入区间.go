package main

import "fmt"

// 57. 插入区间

// 给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
// 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 10}}, []int{2, 5}))
}

// 类似56
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[1] < left { // 在插入区间的左侧且无交集
			res = append(res, interval)
		} else if interval[0] > right { // 在插入区间的右侧且无交集
			if !merged {
				res = append(res, []int{left, right})
				merged = true
			}
			res = append(res, interval)
		} else { // 与插入区间有交集，计算它们的并集
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		res = append(res, []int{left, right})
	}
	return res
}
