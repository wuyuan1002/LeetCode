package main

// LCR 121. 寻找目标值 - 二维数组

// m*n 的二维数组 plants 记录了园林景观的植物排布情况，具有以下特性：
//
// 每行中，每棵植物的右侧相邻植物不矮于该植物；
// 每列中，每棵植物的下侧相邻植物不矮于该植物。
//
// 请判断 plants 中是否存在目标高度值 target。

// findTargetIn2DPlants .
// Offer 04. 二维数组中的查找
func findTargetIn2DPlants(plants [][]int, target int) bool {
	// 从左下角开始找
	i, j := len(plants)-1, 0

	for i >= 0 && j <= len(plants[0])-1 {
		if plants[i][j] > target {
			i--
		} else if plants[i][j] < target {
			j++
		} else {
			return true
		}
	}

	return false
}
