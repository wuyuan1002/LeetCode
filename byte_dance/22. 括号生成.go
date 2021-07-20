package main

// 22. 括号生成

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的括号组合。

func main() {

}

// 回溯法
func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}

	result := make([]string, 0)
	res := make([]byte, 0)
	dfs13(0, 0, n, &res, &result)
	return result
}

func dfs13(left, right, n int, res *[]byte, result *[]string) {
	if left == n && right == n {
		*result = append(*result, string(*res))
		return
	}

	if left < n {
		*res = append(*res, '(')
		dfs13(left+1, right, n, res, result)
		*res = (*res)[:len(*res)-1]
	}
	if left > right {
		*res = append(*res, ')')
		dfs13(left, right+1, n, res, result)
		*res = (*res)[:len(*res)-1]
	}
}
