package main

// 22. 括号生成

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且有效的括号组合。

// func main() {
// 	fmt.Println(generateParenthesis(2))
// }

func generateParenthesis(n int) []string {
	if n < 1 {
		return nil
	}

	result := make([]string, 0)
	res := make([]byte, 0)
	dfs1(0, 0, n, &res, &result)
	return result
}

func dfs1(left, right, n int, res *[]byte, result *[]string) {

	// 若左右括号数量都等于n了，说明得到了一个结果
	if left == n && right == n {
		*result = append(*result, string(*res))
		return
	}

	if left < n {
		*res = append(*res, '(')
		dfs1(left+1, right, n, res, result)
		*res = (*res)[:len(*res)-1]
	}
	if left > right {
		*res = append(*res, ')')
		dfs1(left, right+1, n, res, result)
		*res = (*res)[:len(*res)-1]
	}
}
