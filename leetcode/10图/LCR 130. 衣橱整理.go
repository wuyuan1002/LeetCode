package main

// LCR 130. 衣橱整理

// 家居整理师将待整理衣橱划分为 m x n 的二维矩阵 grid，其中 grid[i][j] 代表一个需要整理的格子。整理师自 grid[0][0] 开始 逐行逐列 地整理每个格子。
//
// 整理规则为：在整理过程中，可以选择 向右移动一格 或 向下移动一格，但不能移动到衣柜之外。同时，不需要整理 digit(i) + digit(j) > cnt 的格子，其中 digit(x) 表示数字 x 的各数位之和。
//
// 请返回整理师 总共需要整理多少个格子。

// wardrobeFinishing .
// Offer 13. 机器人的运动范围
// 模拟机器人运动，从[0, 0]点出发进行递归向下向右移动，若能移动到则说明该作坐标满足要求
func wardrobeFinishing(m int, n int, cnt int) int {
	visited := make([][]bool, m) // 创建地图，默认每个位置都为false，表示没有还走过
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// 返回从[0, 0]点出发所能到达的坐标数量
	return dfsWardrobeFinishing(0, 0, m, n, cnt, visited)
}

// dfsWardrobeFinishing .
// 返回从坐标[i, j]处出发，递归进行向右向下移动，所能到达的坐标位置数量
func dfsWardrobeFinishing(i int, j int, m int, n int, cnt int, visited [][]bool) int {
	// 若当前位置不能走则直接返回0
	if !canMove(i, j, m, n, cnt, visited) {
		return 0
	}

	// 标记当前位置已被访问
	visited[i][j] = true

	// 递归向下向右移动，返回所能到达的位置数量
	return 1 + dfsWardrobeFinishing(i+1, j, m, n, cnt, visited) + dfsWardrobeFinishing(i, j+1, m, n, cnt, visited)
}

// canMove 检查[i,j]位置能不能走，是否符合各位数字只和小于cnt
func canMove(i int, j int, m int, n int, cnt int, visited [][]bool) bool {
	return i >= 0 && i < m && j >= 0 && j < n && !visited[i][j] && num(i)+num(j) <= cnt
}

// num 获取数字n的各位数字只和
func num(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
