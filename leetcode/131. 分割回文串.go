package main

// 131. 分割回文串

// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 回文串 是正着读和反着读都一样的字符串。

func main() {

}

// 回溯法
func partition(s string) [][]string {
	result := make([][]string, 0)
	res := make([]string, 0)
	dfs15(s, 0, &res, &result)
	return result
}

func dfs15(s string, start int, res *[]string, result *[][]string) {
	if start == len(s) {
		tmp := make([]string, len(*res))
		copy(tmp, *res)
		*result = append(*result, tmp)
		return
	}

	for i := start; i < len(s); i++ {
		// 从start位置开始寻找回文串，找到了就放入集合进入下一层递归
		if isPartion(s, start, i) {
			*res = append(*res, s[start:i+1])
			dfs15(s, i+1, res, result)
			*res = (*res)[:len(*res)-1]
		}
	}
}

// 判断是否为回文串
func isPartion(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
