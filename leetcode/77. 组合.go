package main

// 77. 组合

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序返回答案。

func main() {

}

func combine(n int, k int) [][]int {
	if n < k {
		return nil
	}

	res := make([]int, 0)
	result := make([][]int, 0)
	dfs13(1, n, k, &res, &result)
	return result
}

func dfs13(start, n, k int, res *[]int, result *[][]int) {
	if len(*res) == k {
		tmp := make([]int, len(*res))
		copy(tmp, *res)
		*result = append(*result, tmp)
		return
	}

	for i := start; i <= n; i++ {
		*res = append(*res, i)
		dfs13(i+1, n, k, res, result)
		*res = (*res)[:len(*res)-1]
	}
}
