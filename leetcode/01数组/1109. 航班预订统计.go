package main

// 1109. 航班预订统计

// 这里有 n 个航班，它们分别从 1 到 n 进行编号。
//
// 有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi]
// 意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
//
// 请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。

// corpFlightBookings .
// 差分数组
// 换一种思路理解题意，将问题转换为：某公交车共有 n 站，第 i 条记录 bookings[i] = [i, j, k] 表示在 i 站上车 k 人，乘坐到 j 站，在 j+1 站下车，需要按照车站顺序返回每一站车上的人数
// 定义 diff[] 数组记录每站的人数变化，diff[i] 表示第 i 站 -- diff[i]表示第i站相较于前一站的人数变化
// 遍历 bookings[]：bookings[i] = [i, j, k] 表示在 i 站增加 k 人即 diff[i] += k，在 j 站减少 k 人即 diff[j+1] -= k
// 遍历 diff[] 数组，得到每站总人数： 每站的人数为前一站人数加上当前人数变化 diff[i] += diff[i-1]
func corpFlightBookings(bookings [][]int, n int) []int {
	// 创建diff数组用来记录每一站相较于前一站的人数变化 -- 为了计算方便，补充第一站的前一站以及最后一站的后一站方便计算
	diff := make([]int, n+2)

	// 遍历计算每一站的人数变化记录 -- 车站从1到n进行编号
	for _, booking := range bookings {
		// 记录第i站相较于前一站多了多少人，以及第j+1站相较于前一站少了多少人
		diff[booking[0]] += booking[2]
		diff[booking[1]+1] -= booking[2]
	}

	// 遍历diff数组，汇总每一站的实际人数 -- 前一站的总人数加上当前站相对于前一站的人数变化即为当前站的总人数
	for i := 1; i <= n; i++ {
		diff[i] += diff[i-1]
	}

	// 返回第一站到最后一站的总人数
	return diff[1 : n+1]
}

func main() {
	corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5)

}
