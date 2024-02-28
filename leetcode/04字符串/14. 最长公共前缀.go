package main

// 14. 最长公共前缀

// 编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。

// longestCommonPrefix .
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 因为求的是公共前缀，所以公共前缀一定也是第一个字符串的一部分，
	// 因此遍历第一个字符串的所有字符，判断其哪些字符也是其它所有字符串的前缀
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
