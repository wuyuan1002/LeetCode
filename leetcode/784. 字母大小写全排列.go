package main

import "fmt"

// 784. 字母大小写全排列

// 给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。

func main() {
	fmt.Println(letterCasePermutation("a1B2"))
}

func letterCasePermutation(s string) []string {
	if s == "" {
		return nil
	}
	result := make([]string, 0)
	dfs22([]byte(s), 0, &result)
	return result
}

var diff = byte('a' - 'A')

func dfs22(s []byte, index int, result *[]string) {

	// 先把i移动到下一个字母的位置
	i := index
	for ; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' {
			break
		}
	}

	if i == len(s) {
		tmp := make([]byte, len(s))
		copy(tmp, s)
		*result = append(*result, string(tmp))
		return
	}

	if s[i] >= 'a' && s[i] <= 'z' {
		s[i] -= diff
	}
	dfs22(s, i+1, result)

	if s[i] >= 'A' && s[i] <= 'Z' {
		s[i] += diff
	}
	dfs22(s, i+1, result)
}
