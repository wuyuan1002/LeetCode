package main

import "fmt"

// 797. 所有可能的路径

// 给一个有n个结点的有向无环图，找到所有从0到n-1的路径并输出（不要求按顺序）
//
// 二维数组的第 i 个数组中的单元都表示有向图中 i 号结点所能到达的下一些结点
// （译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a ）空就是没有下一个结点了。

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
}

// 类似 省份数量
func allPathsSourceTarget(graph [][]int) [][]int {
	if graph == nil || len(graph) == 0 {
		return nil
	}

	res := make([]int, 0)
	result := make([][]int, 0)

	// 遍历每一个节点
	dfs8(graph, 0, &res, &result)
	return result
}

// 遍历第i个节点所连接的所有节点
func dfs8(graph [][]int, i int, res *[]int, result *[][]int) {
	// 先把当前节点放入路径
	*res = append(*res, i)

	if i == len(graph)-1 {
		tmp := make([]int, len(*res))
		copy(tmp, *res)
		*result = append(*result, tmp)
		return
	}

	// 遍历当前节点连接的所有节点
	for j := 0; j < len(graph[i]); j++ {
		dfs8(graph, graph[i][j], res, result)
		*res = (*res)[:len(*res)-1]
	}
}
