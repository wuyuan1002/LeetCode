package main

// 118. 杨辉三角

// 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。

func main() {

}

// 动态规划
// 类似120
func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}

	dp := make([][]int, numRows)
	for i := range dp {
		dp[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i { // 第一列和最后一列
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			}
		}
	}

	return dp
}
