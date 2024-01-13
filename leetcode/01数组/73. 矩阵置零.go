package main

// 73. 矩阵置零

// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

// setZeroes .
// 两次遍历数组,第一次使用第一行和第一列来标记每一行和每一列是否有零，同时使用两个变量标记第一行和第一列是否原来就有零，第二次置0
func setZeroes(matrix [][]int) {
	// 1. 统计第一行和第一列是否原本就包含0
	row0, col0 := false, false
	for _, n := range matrix[0] {
		if n == 0 {
			row0 = true
			break
		}
	}
	for _, n := range matrix {
		if n[0] == 0 {
			col0 = true
			break
		}
	}

	// 2. 从第二行和第二列开始遍历，查看别的行列是否存在0
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				// 若存在0，则将对应的行和列标记在第一行和第一列上
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 3. 根据第一行第一列的0，将对应的行列置0
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 4. 若第一行第一列原本就有0，则将他们都置为0
	if row0 {
		for j := range matrix[0] {
			matrix[0][j] = 0
		}
	}
	if col0 {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}
