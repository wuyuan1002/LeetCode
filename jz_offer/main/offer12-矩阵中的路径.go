package main

// 矩阵中的路径

// 请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
// 路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。
// 如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
// 例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

// [["a","b","c","e"],
// ["s","f","c","s"],
// ["a","d","e","e"]]

// 但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，
// 路径不能再次进入这个格子。

// func main() {
// 	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
// 	word := "ABCCED"
// 	fmt.Printf("%v", exist1(board, word))
// }

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 || len(board[0]) == 0 || word == "" {
		return false
	}
	rows := len(board)              // 行数
	cols := len(board[0])           // 列数
	visited := make([][]bool, rows) // 存储已经访问过的位置
	for v := range visited {
		visited[v] = make([]bool, cols)
	}

	index := 0 // 当前查找的是第几个字符
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if hasPathCorn(board, i, rows, j, cols, index, word, visited) {
				return true
			}
		}
	}
	return false
}

func hasPathCorn(board [][]byte, row int, rows int, col int, cols int,
	pathLengh int, word string, visited [][]bool) bool {

	if pathLengh >= len(word) {
		return true
	}
	hasPath := false
	if row >= 0 && row < rows && col >= 0 && col < cols &&
		board[row][col] == word[pathLengh] && !visited[row][col] { // 看当前字符是不是

		pathLengh++ // 若当前字符是，则看它的前后左右是不是下一个字符
		visited[row][col] = true
		hasPath = hasPathCorn(board, row+1, rows, col, cols, pathLengh, word, visited) ||
			hasPathCorn(board, row-1, rows, col, cols, pathLengh, word, visited) ||
			hasPathCorn(board, row, rows, col+1, cols, pathLengh, word, visited) ||
			hasPathCorn(board, row, rows, col-1, cols, pathLengh, word, visited)

		if !hasPath { // 若当前字符的前后左右都不是下一个，则当前字符不符合
			visited[row][col] = false
		}
	}
	return hasPath
}

// 第二次做 -- 回溯法
func exist1(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 {
		return false
	}
	if len(word) == 0 {
		return true
	}
	if len(board[0]) == 0 {
		return false
	}

	// 用来记录已经访问过的位置
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := range board {
		for j := range board[i] {
			if hasPathCorn1(board, visited, i, len(board), j, len(board[0]), 0, word) {
				return true
			}
		}
	}
	return false
}

func hasPathCorn1(board [][]byte, visited [][]bool, row, rows, col, cols, index int, word string) bool {

	hasPath := false

	if row >= 0 && row < rows && col >= 0 && col < cols &&
		board[row][col] == word[index] && !visited[row][col] {
		// 标记当前位置已被访问
		visited[row][col] = true
		index++
		if index == len(word) {
			return true
		}

		hasPath = hasPathCorn1(board, visited, row-1, rows, col, cols, index, word) ||
			hasPathCorn1(board, visited, row+1, rows, col, cols, index, word) ||
			hasPathCorn1(board, visited, row, rows, col-1, cols, index, word) ||
			hasPathCorn1(board, visited, row, rows, col+1, cols, index, word)

		if !hasPath {
			visited[row][col] = false
		}
	}

	return hasPath
}
