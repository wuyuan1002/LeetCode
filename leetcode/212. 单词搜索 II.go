package main

import "fmt"

// 212. 单词搜索 II

// 给定一个m x n 二维字符网格board和一个单词（字符串）列表 words，
// 找出所有同时在二维网格和字典中出现的单词。
//
// 单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，
// 其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

func main() {
	board := [][]byte{{'o', 'a', 'b', 'n'}, {'o', 't', 'a', 'e'}, {'a', 'h', 'k', 'r'}, {'a', 'f', 'l', 'v'}}
	words := []string{"oa", "oaa"}
	fmt.Println(findWords(board, words))
}

// 类似79
// 回溯法
func findWords(board [][]byte, words []string) []string {
	if words == nil || len(words) == 0 || board == nil || len(board) == 0 {
		return nil
	}

	res := make([]byte, 0)
	result := make([]string, 0)
	for _, word := range words {
		find := false // 是否找到当前单词
		for i := range board {
			for j := range board[i] {
				if find = dfs20(board, i, j, word, 0, &res); find {
					result = append(result, word)
					break
				}
			}
			if find {
				break
			}
		}
	}

	return result
}

func dfs20(board [][]byte, i, j int, word string, index int, res *[]byte) bool {
	if len(*res) == len(word) {
		return true
	}

	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || board[i][j] != word[index] {
		return false
	}

	char := board[i][j]
	board[i][j] = '*'
	*res = append(*res, char)

	exist := dfs20(board, i-1, j, word, index+1, res) ||
		dfs20(board, i+1, j, word, index+1, res) ||
		dfs20(board, i, j-1, word, index+1, res) ||
		dfs20(board, i, j+1, word, index+1, res)

	*res = (*res)[:len(*res)-1]
	board[i][j] = char
	return exist
}
