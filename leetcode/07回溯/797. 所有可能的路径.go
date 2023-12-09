package main

// 797. 所有可能的路径

// 给一个有n个结点的有向无环图，找到所有从0到n-1的路径并输出（不要求按顺序）
//
// 二维数组的第 i 个数组中的单元都表示有向图中 i 号结点所能到达的下一些结点
// （译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a ）空就是没有下一个结点了。

// allPathsSourceTarget .
// graph中存的是每个节点所能到达的节点列表，所以 graph的长度即len(graph)-1 表示节点数量
func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([]int, 0)      // 一次回溯过程中的结果，一条路径 -- 回溯路径
	result := make([][]int, 0) // 总结果集

	// 因为每一条路径都应该从第0个节点开始，所以先将第0个节点加入结果集，然后进行深度遍历第0个节点所能到达的所有节点
	res = append(res, 0)
	// 开始深度遍历寻找从第0个节出发所能到达的所有节点，统计所有路径
	dfsAllPathsSourceTarget(graph, 0, &res, &result)

	return result
}

// dfsAllPathsSourceTarget .
// 深度优先遍历第i个节点所能到达的下一个节点，当遍历完所有节点说明找到了一条路径，将该路径记录到总结果并进行回溯寻找下一条路径
// graph：选择列表
// i: 遍历寻找从第i个节出发所能到达的所有节点
// res: 一次回溯过程中的结果，一条路径 -- 回溯路径
// result: 总结果集
func dfsAllPathsSourceTarget(graph [][]int, i int, res *[]int, result *[][]int) {
	// 若当前节点已经是最后一个节点，说明已经寻找完了所有节点，已经找到了一条路径，将路径记录到结果集
	// 因为求的是 0~n-1 的节点路径，所以终点一定是最后一个节点（即第 len(graph)-1 个节点），所以这里以当前节点是否为最后一个节点作为判断条件
	// 若题目要求是求从第0个节点出发的所有路径（不要求终点必须是第n-1个节点），那么这里的判断条件需要改成 if len(graph) > 0，即判断从当前节点出发还有没有可到达的其它节点，若没有，说明当前节点是当前路径的最后一个节点，将当前路径记录到结果集
	if i == len(graph)-1 {
		temp := make([]int, len(*res))
		copy(temp, *res)
		*result = append(*result, temp)
		return
	}

	// 遍历第i个节点所能到达的所有节点，将它们分别加入回溯路径并进行回溯
	for j := 0; j < len(graph[i]); j++ {
		// 将当前节点加入到回溯路径
		*res = append(*res, graph[i][j])
		// 递归进入下一层，寻找第j个节点所能到达的所有节点
		dfsAllPathsSourceTarget(graph, graph[i][j], res, result)
		// 将当前节点移出回溯路径
		*res = (*res)[:len(*res)-1]
	}
}
