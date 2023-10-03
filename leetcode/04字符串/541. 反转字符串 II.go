package main

// 541. 反转字符串 II

// 给定一个字符串 s 和一个整数 k，从字符串开头算起，每 2k 个字符反转前 k 个字符。
//
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

// reverseStr .
func reverseStr(s string, k int) string {
	sb := []byte(s)
	// 每2k个字符反转前k个 -- i每次向前移动2k然后反转前k个
	for i := 0; i < len(sb); i += 2 * k {
		// 反转前k个字符，若剩余不足k个，则反转剩余部分
		reverse(sb[i:min(i+k, len(sb))])
	}

	return string(sb)
}

// reverse 反转切片的字符 []
func reverse(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

// min 求最小值
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
