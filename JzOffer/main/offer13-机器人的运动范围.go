package main

// 剑指 Offer 13. 机器人的运动范围

// 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
// 一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），
// 也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，
// 因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

func main() {

}

func movingCount(m int, n int, k int) int {
	if m == 0 || n == 0 || k < 0 {
		return 0
	}

	visited := make([][]bool, m) // 创建地图，默认每个位置都为false，表示没有还走过
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	count := movingCountCorn(0, 0, m, n, k, visited) // 查看下一步能不能走，默认的第一步是[0,0]

	return count

}

// 看 i, j 位置能不能走
func movingCountCorn(i int, j int, m int, n int, k int, visited [][]bool) int {
	count := 0
	if canMove(i, j, m, n, k, visited) { // 检查当前位置能不能走
		visited[i][j] = true
		count = 1 + movingCountCorn(i+1, j, m, n, k, visited) +
			movingCountCorn(i, j+1, m, n, k, visited) +
			movingCountCorn(i-1, j, m, n, k, visited) +
			movingCountCorn(i, j-1, m, n, k, visited)
	}
	return count
}

// 检查[i,j]位置能不能走
func canMove(i int, j int, m int, n int, k int, visited [][]bool) bool {
	return i >= 0 && i < m && j >= 0 && j < n && !visited[i][j] && num(i)+num(j) <= k
}

func num(i int) int {
	sum := 0
	for i > 0 {
		sum += i % 10
		i /= 10
	}
	return sum
}

// 第二次做
func movingCount1(m int, n int, k int) int {
	if m < 0 || n < 0 || k < 0 {
		return -1
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	return moveCount(0, 0, m, n, k, visited)
}

// 计算能移动的方格数
func moveCount(i, j, m, n, k int, visited [][]bool) int {
	count := 0
	if i >= 0 && i < m && j >= 0 && j < n && !visited[i][j] && num1(i)+num1(j) <= k {
		visited[i][j] = true
		count = 1 + moveCount(i+1, j, m, n, k, visited) +
			moveCount(i-1, j, m, n, k, visited) +
			moveCount(i, j+1, m, n, k, visited) +
			moveCount(i, j-1, m, n, k, visited)
	}
	return count
}

func num1(i int) int {
	num := 0
	for i > 0 {
		num += i % 10
		i /= 10
	}
	return num
}
