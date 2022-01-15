package main

// 79. 单词搜索

// 给定一个m x n 二维字符网格board 和一个字符串单词word 。
// 如果word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，
// 其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

// func main() {
// 	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
// 	word := "ABCCED"
// 	fmt.Printf("%v", exist(board, word))
// }

// 矩阵中的路径, offer 12
// 回溯法
func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	if word == "" {
		return true
	}

	// 路径
	res := make([]byte, 0, len(word))

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs7(board, i, j, word, 0, &res) {
				return true
			}
		}
	}
	return false
}

func dfs7(board [][]byte, i, j int, word string, index int, res *[]byte) bool {
	if len(*res) == len(word) {
		return true
	}

	exist := false

	// 若当前位置可以访问
	if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) &&
		board[i][j] == word[index] && board[i][j] != '*' {

		// 记录下当前字符并把当前位hi标记为已访问
		char := board[i][j]
		board[i][j] = '*'

		// 将当前字符加入到路径中
		*res = append(*res, char)
		// 探索前后左右位置
		exist = dfs7(board, i+1, j, word, index+1, res) ||
			dfs7(board, i-1, j, word, index+1, res) ||
			dfs7(board, i, j+1, word, index+1, res) ||
			dfs7(board, i, j-1, word, index+1, res)

		// 将当前字符移出路径，并标记当前位置为未访问
		*res = (*res)[:len(*res)-1]
		board[i][j] = char
	}

	return exist
}
