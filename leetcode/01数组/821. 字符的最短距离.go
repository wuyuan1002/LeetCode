package main

import "math"

// 821. 字符的最短距离

// 给你一个字符串 s 和一个字符 c ，且 c 是 s 中出现过的字符。
//
// 返回一个整数数组 answer ，其中 answer.length == s.length 且 answer[i] 是 s 中从下标 i 到离它 最近 的字符 c 的 距离 。
//
// 两个下标 i 和 j 之间的 距离 为 abs(i - j) ，其中 abs 是绝对值函数。

// shortestToChar .
// 两次遍历字符串，第一次正序遍历计算每个字符与其左侧最近c的距离，第二次倒序遍历计算每个字符与其右侧最近c的距离，两者取最小值
func shortestToChar(s string, c byte) []int {
	result := make([]int, len(s))

	// 将每个位置的初始距离设置为最大值，表示当前字符左右侧的c还没出现
	for i := range result {
		result[i] = math.MaxInt
	}

	// 第一次正序遍历字符串，计算出每个字符与其左侧最近c的距离
	for i, idx := 0, -1; i < len(s); i++ {
		// 若当前字符本身为c，则记录当前c出现的下标
		if s[i] == c {
			idx = i
		}
		// 计算出当前字符与其左侧最近c的距离
		if idx != -1 {
			result[i] = i - idx
		}
	}

	// 第二次倒序遍历字符串，计算每个字符与其右侧最近c的距离，并与左侧c的距离取最小值
	for i, idx := len(s)-1, -1; i >= 0; i-- {
		// 若当前字符本身为c，则记录当前c出现的下标
		if s[i] == c {
			idx = i
		}
		// 计算当前字符与其左右侧c的距离的最小值 -- min(当前字符与左侧最近的c的距离, 当前字符与右侧最近的c的距离)
		if idx != -1 {
			result[i] = min(result[i], idx-i)
		}
	}

	return result
}
