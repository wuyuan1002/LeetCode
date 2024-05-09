package main

// 1790. 仅执行一次字符串交换能否使两个字符串相等

// 给你长度相等的两个字符串 s1 和 s2 。一次 字符串交换 操作的步骤如下：选出某个字符串中的两个下标（不必不同），并交换这两个下标所对应的字符。
//
// 如果对 其中一个字符串 执行 最多一次字符串交换 就可以使两个字符串相等，返回 true ；否则，返回 false 。

// areAlmostEqual .
// 只执行一次交换使字符串相等，意味着两个字符串中最多只存在两个位置 a、b 处字符不相等，此时我们交换 a、b 处字符可使其相等
func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	} else if s1 == s2 {
		return true
	}

	// 使用两个下标记录两字符串不同字符的位置
	a, b := -1, -1
	for i := range s1 {
		if s1[i] == s2[i] {
			continue
		}

		if a < 0 {
			a = i
		} else if b < 0 {
			b = i
		} else {
			return false
		}
	}

	// 若两位置字符交叉相等说明可以交换使两字符串相等
	return s1[a] == s2[b] && s1[b] == s2[a]
}
