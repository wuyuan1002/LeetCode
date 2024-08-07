package main

import "sort"

// 455. 分发饼干

// 假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
//
// 对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；
// 并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，
// 这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

// findContentChildren .
// 优先使用大饼干满足胃口大的小孩，即用最合适的饼干满足每一个小孩
//
// 为了满足更多的小孩，就不要造成饼干尺寸的浪费
// 大尺寸的饼干既可以满足胃口大的孩子也可以满足胃口小的孩子，那么就应该优先满足胃口大的
// 这里的局部最优就是大饼干喂给胃口大的，充分利用饼干尺寸喂饱一个，全局最优就是喂饱尽可能多的小孩
//
// 使用贪心策略，先将饼干数组和小孩数组排序，然后从后向前遍历小孩数组，用大饼干优先满足胃口大的，并统计满足的小孩数量
func findContentChildren(g []int, s []int) int {
	// 按照小孩胃口和饼干大小分别升序排序小孩数组和饼干数组
	sort.Ints(g)
	sort.Ints(s)

	result := 0         // 总结果 -- 能被满足的小孩个数
	index := len(s) - 1 // 饼干数组下标 -- 优先使用大饼干满足胃口大的小孩 -- 倒序遍历饼干数组

	// 遍历小孩数组，优先满足胃口大的小孩 -- 倒序遍历小孩数组
	for i := len(g) - 1; i >= 0; i-- {
		// 若还有饼干且当前饼干能够满足当前小孩 -- 结果值+1，并继续遍历下一个尺寸较小的饼干
		if index >= 0 && s[index] >= g[i] {
			result++
			index--
		}
	}

	return result
}
