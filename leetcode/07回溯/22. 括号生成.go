package main

// 22. 括号生成

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

// generateParenthesis .
//
// 递归回溯，先放左括号，再放右括号，当左右括号各数都等于n时即表示得到一个结果
func generateParenthesis(n int) []string {
	res := make([]byte, 0, 2*n)
	result := make([]string, 0, n)
	dfsGenerateParenthesis(0, 0, n, &res, &result)
	return result
}

// dfsGenerateParenthesis .
//
// 往回溯列表中先放左括号, 再放右括号, 当左右括号各数都等于n时表示得到一个结果，将结果保存到结果集中
//
// left: 左括号已放个数
// right: 右括号已放个数
// n: 总共放几对括号
// res: 回溯路径 -- 保存一次回溯所放的括号
// result: 结果集 -- 存放满足条件的括号组合
func dfsGenerateParenthesis(left, right, n int, res *[]byte, result *[]string) {
	// 当左右括号都放了n个，表示得到了一个n对括号的组合，将结果存放到结果集
	if left == n && right == n {
		*result = append(*result, string(*res))
		return
	}

	// 优先放左括号 -- 因为需要先有左括号才能放右括号来与其对应
	if left < n {
		*res = append(*res, '(')
		dfsGenerateParenthesis(left+1, right, n, res, result)
		*res = (*res)[:len(*res)-1]
	}

	// 其次放右括号 -- 为什么不用(right < n)这个条件呢？
	// 因为只有左括号数量大于右括号数量时, 放右括号才是合法的, 才会有正确的左括号匹配
	//
	// 这里主要是防止前一步从回溯列表移出一个左括号后，会继续向下执行到此添加右括号，
	// 如果移出左括号后整体的左括号数量不大于右括号数量（比如移出的是最后一个左括号），
	// 说明回溯列表中没有多余的左括号可以与接下来要添加的右括号配对，所以此时不能添加右括号
	if right < left {
		*res = append(*res, ')')
		dfsGenerateParenthesis(left, right+1, n, res, result)
		*res = (*res)[:len(*res)-1]
	}
}
