package main

// 22. 括号生成

// func main() {

// }

func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}

	res := make([]byte, 0, 2*n)
	result := make([]string, 0, n)
	dfs10(0, 0, n, &res, &result)
	return result
}

func dfs10(left, right, n int, res *[]byte, result *[]string) {
	if left == n && right == n {
		*result = append(*result, string(*res))
		return
	}

	if left < n {
		*res = append(*res, '(')
		dfs10(left+1, right, n, res, result)
		*res = (*res)[:len(*res)-1]
	}
	if right < left {
		*res = append(*res, ')')
		dfs10(left, right+1, n, res, result)
		*res = (*res)[:len(*res)-1]
	}
}
