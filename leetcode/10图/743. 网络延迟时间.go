package main

import "math"

// 743. 网络延迟时间

// 有 n 个网络节点，标记为 1 到 n。
//
// 给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， wi 是一个信号从源节点传递到目标节点的时间。
//
// 现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。

// networkDelayTime .
// Dijkstra 算法
// 题目的目的是 求从节点k出发，所有点都被访问到的最短路径，也就是 -> 从节点k出发到达其它所有节点的最短路径的最大值，若存在从k出发无法到达的点，则返回-1
//
// 将所有节点分为两类，一类为已确定好的节点，即从节点k出发到该节点的最短路径已经计算完成，一类为未确定节点，即从节点k出发到该节点的最短路径还未计算完成
//
// 定义 g[a][b] = c 表示节点a到节点b的边的权重为c，若a到b没有边，则c=MaxInt -- 使用邻接矩阵来进行存图
// 定义 dist[a] 表示从节点k到节点a的最短路径，初始值dis[k]=0表示节点k到其自身的路径长度为0，其余dis[b]=MaxInt表示节点k到节点b的最短路径长度尚未计算出 -- 在计算过程中不断更新该值直到求到最小值
// 定义 done[a] 表示从节点k到节点a的最短路径已经计算完成，即dist[a]已经不能再继续变小了 -- done[a]=true表示节点a为已确定节点，done[b]=false表示节点b为未确定节点
func networkDelayTime(times [][]int, n int, k int) int {
	// 定义无效值 -- 设置为MaxInt/2而不是MaxInt的原因是防止后面在加法运算时溢出
	const INF = math.MaxInt / 2

	// 邻接矩阵用来存图
	g := make([][]int, n)
	// 初始化邻接矩阵，先将每个节点间的路径设置为MaxInt表示当前两节点间还不存在边
	for a := range g {
		g[a] = make([]int, n)
		for b := range g[a] {
			// 将任意两节点间的边的权重设置为INF表示当前两节点间还没有边
			g[a][b] = INF
		}
	}
	// 初始化图信息，将节点间边的信息补充到邻接矩阵中 -- 邻接矩阵中节点编号从0开始
	for _, t := range times {
		g[t[0]-1][t[1]-1] = t[2]
	}

	// -------

	// 定义dist用来存储节点k到每一个节点的最短路径长度，初始设置为INF表示最短路径还没计算出
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	// 设置节点k到其自身的最短路径长度为0
	dist[k-1] = 0

	// -------

	// done[a]表示从节点k到节点a的最短路径已经计算完成，即dist[a]不能再变的更小 -- 初始值默认都为false
	done := make([]bool, n)

	// -------

	// 遍历每一个节点i，计算从节点k到该节点i的最短路径长度
	for i := 0; i < n; i++ {
		// 不断从 未确定节点 中取一个与起点k距离最短的点x，将它归类为 已确定节点，并用它更新从起点到其他所有 未确定节点 的距离。直到所有点都被归类为 已确定节点
		x := -1
		for y, isDone := range done {
			if !isDone && (x == -1 || dist[y] < dist[x]) {
				x = y
			}
		}

		// 标记节点x为已确定节点 -- 即节点k到节点x的最短路径长度已计算完成，dist[x]已经是最小值了
		done[x] = true

		// 用节点k到节点x的最短路径去更新每个节点z的最短路径长度
		// 用节点x更新节点z的意思是，用起点k到节点x的最短路长度加上从节点x到节点z的边的长度，去比较原来计算的从起点k到节点z的最短路长度，如果前者小于后者，就用前者更新后者
		for z, time := range g[x] {
			dist[z] = min(dist[z], dist[x]+time)
		}
	}

	// -------

	// 每个节点都计算完后，节点k到每个节点的最短路径已经计算结束，即dist[a]，接下来计算这些最短路径长度中的最大值即为答案
	result := -1
	for _, d := range dist {
		if d == INF {
			// 若存在节点的最短路径长度为INF说明从节点k无法到达该节点，直接返回-1
			return -1
		}
		result = max(result, d)
	}

	return result
}

// min .
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
