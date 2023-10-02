package main

// 541. 反转字符串 II

// 给定一个字符串 s 和一个整数 k，从字符串开头算起，每 2k 个字符反转前 k 个字符。
//
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

// func main() {

// }

func reverseStr(s string, k int) string {
	reverse := func(s []byte) {
		l, r := 0, len(s)-1
		for l < r {
			s[l], s[r] = s[r], s[l]
			l++
			r--
		}
	}

	sb := []byte(s)
	for i := 0; i < len(sb); i += 2 * k {
		if i+k < len(sb) {
			reverse(sb[i : i+k])
		} else {
			reverse(sb[i:])
		}
	}

	return string(sb)
}
