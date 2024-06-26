package main

// 785. 判断二分图

// 存在一个 无向图 ，图中有 n 个节点。其中每个节点都有一个介于 0 到 n - 1 之间的唯一编号。
// 给你一个二维数组 graph ，其中 graph[u] 是一个节点数组，由节点 u 的邻接节点组成。
// 形式上，对于 graph[u] 中的每个 v ，都存在一条位于节点 u 和节点 v 之间的无向边。该无向图同时具有以下属性：
// 不存在自环（graph[u] 不包含 u）。
// 不存在平行边（graph[u] 不包含重复值）。
// 如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图）
// 这个图可能不是连通图，也就是说两个节点 u 和 v 之间可能不存在一条连通彼此的路径。
// 二分图 定义：如果能将一个图的节点集合分割成两个独立的子集 A 和 B ，并使图中的每一条边的两个节点一个来自 A 集合，一个来自 B 集合，就将这个图称为 二分图 。
//
// 如果图是二分图，返回 true ；否则，返回 false 。

// isBipartite .
// 也就是说能否将图中的节点分成两个集合，使得图中的任意一条边所连接的两个节点都分别属于不同的集合，即不存在一条边上的两节点同属于一个集合的情况
//
// 1. 深度优先遍历 / 广度优先遍历
// 	从任意节点开始遍历图，使用两种颜色进行染色，将连接的节点分别染成相反的颜色，连接过程中若发现相邻顶点被染成了相同的颜色则说明不是二分图，否则说明是二分图
// 2. 并查集
// 	根据二分图的定义可知，二分图的每个节点的所有邻接点都应该属于同一集合，且不与该节点属于同一集合
// 	使用并查集处理元素分组问题，遍历图中的每个节点，将其所有邻接点合并到同一集合，若发现这些邻接点中有节点已经与当前节点属于同一集合则说明不是二分图
func isBipartite(graph [][]int) bool {
	// 并查集 -- parent[i]表示i号节点，每个节点编号为[0, n-1]
	parent := make([]int, len(graph))
	for i := range parent {
		parent[i] = i
	}

	// 遍历每个节点，将其所有邻接点进行合并（合并到同一集合），过程中判断是否违背了二分图的定义
	for i := range graph {
		// 遍历节点i的所有邻接点并将其进行合并
		for _, n := range graph[i] {
			// 若节点i的某个邻接点n已经与节点i属于同一集合则说明不是二分图，直接返回即可
			if isConnect(parent, i, n) {
				return false
			}
			// 将当前邻接点n与节点i的第一个邻接点连接，即与节点i的第一个邻接点合并到同一集合中
			connect(parent, graph[i][0], n)
		}
	}

	return true
}
