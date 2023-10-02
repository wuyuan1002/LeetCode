package main

// 119. 杨辉三角 II

// 给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。

// func main() {
// 	fmt.Println(getRow(3))
// }

// 类似118 -- 只是本题可以使用一维数组滚动优化
func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}

	dp := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ { // 遍历第几行
		for j := i; j >= 0; j-- { // 遍历当前行
			if j == 0 || j == i {
				dp[j] = 1
			} else {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp
}
