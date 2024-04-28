package main

// 剑指 Offer 04. 二维数组中的查找

// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，
// 每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，
// 输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

// func main() {
// 	// arr := [][]int{{1, -1}}
// 	// arr := [][]int{{-5}}
// 	arr := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}, {19, 23, 25, 34, 35}}
// 	fmt.Print(findNumberIn2DArray2(arr, 5))
// }

// 同Hot100 240
func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 从左下角开始找

	lenLie := len(matrix) // 某一列有几个元素
	var lenHang int       // 某一行有几个元素
	if lenLie != 0 {
		lenHang = len(matrix[0]) // 二维数组每行的长度 5
	} else {
		lenHang = 0
	}

	i := lenLie - 1 // 遍历某一列
	j := 0          // 遍历某一行

	for i >= 0 && j < lenHang {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}

// 第二次做
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	// 从左下角开始找

	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	maxLieIndex := len(matrix) - 1     // 一列的最大下标
	maxHangIndex := len(matrix[0]) - 1 // 一行的最大下标

	i, j := maxLieIndex, 0 // i第几行，j第几列
	for i >= 0 && j <= maxHangIndex {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}

func findNumberIn2DArray3(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}

	lieNum := len(matrix) - 1     // 列长度
	hangNum := len(matrix[0]) - 1 // 行长度

	i, j := lieNum, 0 // 左下角开始遍历
	for i >= 0 && j <= hangNum {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}
