package main

// 14. 最长公共前缀

// 编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。

// longestCommonPrefix .
//
// 因为求的是所有字符串的公共前缀，所以公共前缀一定也是第一个字符串的一部分，
// 因此遍历第一个字符串的所有字符，判断其哪些字符也是其它所有字符串的前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 从头到尾遍历第一个字符串的的每个字符
	for i := 0; i < len(strs[0]); i++ {
		// 不断遍历第j个字符串，判断其第i个位置的字符是否与第一个字符串一致
		for j := 1; j < len(strs); j++ {
			// 若当前字符未出现在当前字符串的相同位置，说明最长公共前缀就是到当前字符前一个字符的子串
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	// 若第一个字符串的所有字符都出现在所有字符串的相同位置，那么第一个字符串整体就是最长公共前缀
	return strs[0]
}
