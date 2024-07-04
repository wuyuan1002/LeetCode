package main

// 36. 有效的数独

// 请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
//
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
//
//
// 注意：
//
// 一个有效的数独（部分已被填充）不一定是可解的。
// 只需要根据以上规则，验证已经填入的数字是否有效即可。
// 空白格用 '.' 表示。

// isValidSudoku .
// 使用多维数组或map存储每行每列每块的数字出现情况
func isValidSudoku(board [][]byte) bool {
	if board == nil || len(board) == 0 {
		return false
	}

	// 记录每行每列每块中每个数字的出现次数
	// - row[i][num]表示第i行数字num的出现次数
	// - col[j][num]表示第j列数字num的出现次数
	// - box[i][j][num]表示第[i/3, j/3]个小九宫格中数字num的出现次数 -- 数字board[i][j]所属的九宫格为box[i/3, j/3]
	row, col, box := [9][9]int{}, [9][9]int{}, [3][3][9]int{}

	// 遍历数独的每个数字，记录每个行、列、九宫格中数字出现的次数，并判断数字出现的次数是否符合数独规则
	for i := range board { // 遍历每一行
		for j := range board[i] { // 遍历每一列
			if board[i][j] != '.' {
				num := board[i][j] - '1'

				// 将当前数字出现的次数更新到每行每列每块中
				row[i][num]++
				col[j][num]++
				box[i/3][j/3][num]++

				// 若当前数字在行列块中出现次数大于1，则不是有效的数独
				if row[i][num] > 1 || col[j][num] > 1 || box[i/3][j/3][num] > 1 {
					return false
				}
			}
		}
	}

	return true
}
