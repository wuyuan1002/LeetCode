package main

import "sort"

// 406. 根据身高重建队列

// 假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。
// 每个 people[i] = [hi, ki] 表示第 i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。
//
// 请你重新构造并返回输入数组 people 所表示的队列。返回的队列应该格式化为数组 queue ，
// 其中 queue[j] = [hj, kj] 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。

// reconstructQueue .
// 给定两个维度，一个是身高h，一个是前面比自己高的人的个数k
// 先排序，再插队：
// - 将队列按照两个维度进行排序，先按照h升序排序，h相等的则按照k降序排序
// - 从头到尾遍历排序数组进行插入，将每个元素插入到数组的下标为k的位置（表示该元素前面有k个比自己高或相同高度的人数）
func reconstructQueue(people [][]int) [][]int {
	// 将队列按照两个维度进行排序，先按照h升序排序，h相等的则按照k降序排序
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})

	// 从头到尾遍历排序数组进行插入，将每个元素插入到数组的下标为k的位置（表示该元素前面有k个比自己高或相同高度的人数）
	result := make([][]int, 0, len(people))
	for _, person := range people {
		idx := person[1]
		result = append(result[:idx], append([][]int{person}, result[idx:]...)...)
	}

	return result
}
