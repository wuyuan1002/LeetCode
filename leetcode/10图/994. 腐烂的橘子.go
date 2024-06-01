package main

// 994. 腐烂的橘子

// 在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
//
// 值 0 代表空单元格；
// 值 1 代表新鲜橘子；
// 值 2 代表腐烂的橘子。
// 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
//
// 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。

// orangesRotting .
// BFS
// 1. 首先遍历一遍整个网格，统计出新鲜橘子的数量，记为count，并且将所有腐烂的橘子的坐标加入队列queue中
// 2. 接下来进行广度优先搜索，每一轮搜索，将队列中的所有腐烂的橘子向四个方向腐烂新鲜橘子，直到队列为空或者新鲜橘子的数量为0为止
// 3. 最后，如果新鲜橘子的数量为0，则返回当前的轮数，否则返回-1
func orangesRotting(grid [][]int) int {
	result := 0                // 总结果 -- 所有橘子腐烂所需的时间
	count := 0                 // 新鲜橘子的数量
	queue := make([][2]int, 0) // 存放腐烂的橘子

	// 1.
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				count++
			} else if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	// 2.
	for len(queue) > 0 && count > 0 {
		// 存放当前分钟将被腐烂的橘子，也就是下一分钟将向四周扩散的节点 -- 即下一层节点
		q := make([][2]int, 0)

		// 每分钟将当前队列中的所有腐烂橘子向四周扩散一次，将新腐烂的橘子加入到队列中，用来下一分钟向四周扩散
		for i := range queue { // 遍历每一个腐烂的橘子
			for _, dir := range dirs { // 向四周扩散
				row, col := queue[i][0]+dir[0], queue[i][1]+dir[1]
				if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0]) && grid[row][col] == 1 {
					// 将新腐烂的橘子加入到新的队列中
					q = append(q, [2]int{row, col})
					// 标记当前位置的橘子为腐烂
					grid[row][col] = 2
					// 减少新鲜橘子的个数
					count--
				}
			}
		}

		// 消耗了一分钟，增加耗时、同时将队列替换为下一分钟要扩散的队列
		result++
		queue = q
	}

	// 3.
	if count > 0 {
		return -1
	}
	return result
}
