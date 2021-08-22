package main

// 216. 组合总和 III

// 找出所有相加之和为n 的k个数的组合。组合中只允许含有 1 -9 的正整数，并且每种组合中不存在重复的数字。
//
// 说明：
// 所有数字都是正整数。
// 解集不能包含重复的组合。

func main() {

}

func combinationSum4(k int, n int) [][]int {
	result := make([][]int, 0)
	res := make([]int, 0)
	dfs14(1, k, 0, n, &res, &result)

	return result
}

func dfs14(start, k, currentSum, targetSum int, res *[]int, result *[][]int) {
	if currentSum == targetSum && len(*res) == k {
		tmp := make([]int, len(*res))
		copy(tmp, *res)
		*result = append(*result, tmp)
		return
	}

	for i := start; i <= 9; i++ {
		sum := currentSum + i
		if sum > targetSum || len(*res) == k {
			// 剪枝
			return
		}

		*res = append(*res, i)
		dfs14(i+1, k, sum, targetSum, res, result)
		*res = (*res)[:len(*res)-1]
	}
}
