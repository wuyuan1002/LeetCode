package main

// 79. 单词搜索

// 给定一个m x n 二维字符网格board 和一个字符串单词word 。
// 如果word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，
// 其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

// exist .
// 遍历地图的每个位置，判断从该位置出发能否找到指定单词，若能找到则直接返回
func exist79(board [][]byte, word string) bool {
	// 回溯路径
	res := make([]byte, 0, len(word))

	// 遍历从地图的每个位置出发，能否找到单词
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfsExist(board, word, i, j, 0, &res) {
				return true
			}
		}
	}
	return false
}

// dfsExist 返回从给地图定位置出发是否能够找到index之后的字符串
// board: 地图
// word: 整个字符串
// i, j: 地图出发的下标
// index: 从指定下标出发寻找index之后的字符串
// res: 回溯路径
func dfsExist(board [][]byte, word string, i, j int, index int, res *[]byte) bool {
	// 整个字符串已经找到，直接返回true
	if len(*res) == len(word) {
		return true
	}

	// 从[i, j]位置出发是否能够找到index之后的字符串
	exist := false

	// 若当前位置在地图范围内，且当前位置字符为第index个字符且没有访问过，说明在当前位置找到了一个字母，开始寻找字符串的下一个字母
	if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) && board[i][j] == word[index] && board[i][j] != '*' {
		// 将当前字母加入到路径中
		*res = append(*res, word[index])

		// 标记当前位置已被访问过
		board[i][j] = '*'

		// 从当前位置的上下左右寻找能不能找到字符串的下一个字母
		exist = dfsExist(board, word, i+1, j, index+1, res) ||
			dfsExist(board, word, i-1, j, index+1, res) ||
			dfsExist(board, word, i, j+1, index+1, res) ||
			dfsExist(board, word, i, j-1, index+1, res)

		// 将当前字符移出回溯路径，并标记当前位置为未访问
		*res = (*res)[:len(*res)-1]
		board[i][j] = word[index]
	}

	return exist
}
