package main

// 14. 最长公共前缀

// 编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。

func main() {

}

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	prefix := ""        // 公共前缀
	for i := 0; ; i++ { // 对每个字符串比较第i个字符是否相等
		str := strs[0]
		if i > len(str)-1 {
			break
		}
		tmp := str[i]
		j := 1
		for ; j < len(strs); j++ { // 挨个比较每个字符串的第i个字符
			str = strs[j]
			if i > len(str)-1 {
				break
			}
			if str[i] != tmp {
				break
			}
		}

		if j != len(strs) {
			// 有字符串长度不够或第i个字符不相等
			break
		}
		prefix += string(tmp)
	}

	return prefix
}
