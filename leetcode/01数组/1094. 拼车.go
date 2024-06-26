package main

// 1094. 拼车

// 车上最初有 capacity 个空座位。车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向）
//
// 给定整数 capacity 和一个数组 trips , trip[i] = [numPassengersi, fromi, toi] 表示第 i 次旅行有 numPassengersi 乘客，
// 接他们和放他们的位置分别是 fromi 和 toi 。这些位置是从汽车的初始位置向东的公里数。
//
// 当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false。

// carPooling .
// 同 leetcode 1109. 航班预订统计
// 查分数组
//
// 定义 diff[] 数组记录每站的人数变化，diff[i] 表示第 i 站 -- diff[i]表示第i站相较于前一站的人数变化
// 遍历 trips[]：trips[i] = [i, j, k] 表示在 i 站增加 k 人即 diff[i] += k，在 j 站减少 k 人即 diff[j+1] -= k
// 遍历 diff[] 数组，得到每站总人数： 每站的人数为前一站人数加上当前人数变化 diff[i] += diff[i-1]
//
// 举例 -- carPooling([][]int{{10, 0, 2}, {20, 2, 3}, {25, 2, 5}}, 50)
// 计算各站人数变化（车站编码从0开始，车站从0到n进行编号，trip中记录的from和to是左闭右开的，即实际乘坐的范围是[from, to)）：
// +-----------------------------------+
// |  0  |  1  |  2  |  3  |  4  |  5  |
// |     |     |     |     |     |     |
// | +10 |     | -10 |     |     |     |
// |     |     | +20 | -20 |     |     |
// |     |     | +25 |     |     | -25 |
// +-----------------------------------+
// | +10 |  0  | +35 | -20 |  0  | -25 |  -- 汇总各站人数总变化
// +-----------------------------------+
// 统计车上总人数：
// 0: 0+10=10
// 1: 10+0=10
// 2: 10+35=45
// 3: 45-20=25
// 4: 25+0=25
// 5: 25-25=0
func carPooling(trips [][]int, capacity int) bool {
	// 获取车能到达的最远的站 -- 车站总个数为 farthest+1（要包含第0站），即[0, farthest]
	farthest := 0
	for _, trip := range trips {
		farthest = max(farthest, trip[2])
	}

	// 创建diff数组用来记录每一站相较于前一站的人数变化 -- 因为车站记录的就是上车和下车的准确站数，即实际乘坐的范围是[from, to)，所以不必补充最后一站的后一站以及第一站的前一站
	diff := make([]int, farthest+1)

	// 遍历计算每一站的人数变化记录 -- 车站编码从0开始，车站从0到n进行编号
	for _, trip := range trips {
		// 记录第from站相较于前一站多了多少人，以及第to站相较于前一站少了多少人 -- trip中记录的from和to是左闭右开的，即实际乘坐的范围是[from, to)
		diff[trip[1]] += trip[0]
		diff[trip[2]] -= trip[0]
	}

	// 遍历diff数组，汇总每一站的实际人数，若总人数大于车的最大容量则说明无法接送所有乘客 -- 前一站的总人数加上当前站相对于前一站的人数变化即为当前站的总人数
	for i := 0; i <= farthest; i++ {
		// 第一站的人数就是车上的实际人数，无需考虑与前一站的人数变化 -- 这里因为车站编号是从0开始的，所以没办法在前面增加虚拟的前一站
		if i > 0 {
			diff[i] += diff[i-1]
		}

		// 若当前车上总人数已超过车的容量，则说明无法接送所有乘客
		if diff[i] > capacity {
			return false
		}
	}

	return true
}
