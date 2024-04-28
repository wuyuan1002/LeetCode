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
// 	fmt.Printf("%v", exist(board, word))
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

// 第二次做
func exist1(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	} else if len(board) == 0 {
		return false
	}

	rows := len(board)    // 行数
	cols := len(board[0]) // 列数

	// visited用于记录对应位置是否被访问过
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if havePath(board, i, rows, j, cols, 0, word, visited) {
				return true
			}
		}
	}
	return false
}

func havePath(board [][]byte, i int, rows int, j int, cols int, index int, word string, visited [][]bool) bool {
	if index >= len(word) {
		return true
	}

	have := false
	if i >= 0 && i < rows && j >= 0 && j < cols && board[i][j] == word[index] && !visited[i][j] {
		visited[i][j] = true // 标记当前位置已被访问

		have = havePath(board, i+1, rows, j, cols, index+1, word, visited) ||
			havePath(board, i-1, rows, j, cols, index+1, word, visited) ||
			havePath(board, i, rows, j+1, cols, index+1, word, visited) ||
			havePath(board, i, rows, j-1, cols, index+1, word, visited)

		if !have {
			visited[i][j] = false // 将当前位置标记为未访问
		}
	}
	return have
}
