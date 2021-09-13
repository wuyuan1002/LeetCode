package main

// 447. 回旋镖的数量

// 给定平面上n 对 互不相同 的点points ，其中 points[i] = [xi, yi] 。
// 回旋镖是由点(i, j, k) 表示的元组 ，其中i和j之间的距离和i和k之间的距离相等（需要考虑元组的顺序）。
//
// 返回平面上所有回旋镖的数量。

func main() {

}

// 1. 遍历每一个点，使用map保存剩余点到当前点的距离
// 2. 回溯法 -- 不断选择三个点计算距离，将距离相等的放如结果集
func numberOfBoomerangs(points [][]int) int {
	if points == nil || len(points) < 3 {
		return 0
	}

	res := 0
	for _, p := range points { // 遍历每一个点
		hash := make(map[int]int) // 保存距离和对应的点的个数
		for _, q := range points {
			// 计算距离
			dis := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])

			// 将距离存入map中
			hash[dis]++
		}

		// 遍历每一个距离，相加总结果
		for _, disnum := range hash {
			res += disnum * (disnum - 1)
		}
	}

	return res
}
