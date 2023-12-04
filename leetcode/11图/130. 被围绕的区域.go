package main

// 130. 被围绕的区域

// 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，
// 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

// solve .
// 同 leetcode 1020
// 相当于把所有的飞地（四周都是水）都用水填充
func solve(board [][]byte) {
	// 遍历地图的每一个坐标，将每个飞地用水覆盖
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// 若当前节点已被访问过则为A、水为X、只有当前节点是陆地且未被访问过为O
			if board[i][j] == 'O' {
				// 首先判断当前坐标所属陆地是否为飞地，同时将陆地标记为A
				isFlyArea := true
				dfsSolve(board, i, j, &isFlyArea, 'O', 'A')
				if isFlyArea {
					// 若当前陆地是飞地，则将当前陆地用水覆盖，即将当前陆地A用X覆盖
					dfsSolve(board, i, j, &isFlyArea, 'A', 'X')
				} else {
					// 若当前陆地不是飞地，则将当前陆地恢复为原来的样子，即将现在的A用原来的O覆盖
					dfsSolve(board, i, j, &isFlyArea, 'A', 'O')
				}
			}
		}
	}
}

// dfsSolve 判断给定坐标所属陆地是否为飞地，同时将陆地坐标标记为给定标记
func dfsSolve(board [][]byte, i, j int, isFlyArea *bool, base, mark byte) {
	// 若当前坐标超出地图范围或不是陆地或已被访问过
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != base {
		return
	}

	// 标记当前位置已被访问
	board[i][j] = mark

	// 若当前坐标不在地图边界（是飞地）、若当前坐标为地图边界（不是飞地）
	// 若isFlyArea已经是false，说明当前坐标所属陆地已经不是飞地了，不应重复赋值防止覆盖 -- 所以前面要加上&&（false和任何数做与运算都为false）
	*isFlyArea = *isFlyArea && !(i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1)

	// 上下左右继续遍历
	dfsSolve(board, i+1, j, isFlyArea, base, mark)
	dfsSolve(board, i-1, j, isFlyArea, base, mark)
	dfsSolve(board, i, j+1, isFlyArea, base, mark)
	dfsSolve(board, i, j-1, isFlyArea, base, mark)
}
