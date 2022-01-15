package main

// 1221. 分割平衡字符串

// 在一个 平衡字符串 中，'L' 和 'R' 字符的数量是相同的。
// 给你一个平衡字符串s，请你将它分割成尽可能多的平衡字符串。
//
// 注意：分割得到的每个字符串都必须是平衡字符串。
// 返回可以通过分割得到的平衡字符串的 最大数量 。

// func main() {

// }

func balancedStringSplit(s string) int {
	res := 0
	nl := 0 // L的个数
	for _, c := range s {
		if c == 'L' {
			nl++
		} else if c == 'R' {
			nl--
		}

		if nl == 0 {
			res++
		}
	}

	return res
}
