package main

// 409. 最长回文串

// 给定一个包含大写字母和小写字母的字符串 s ，返回 通过这些字母构造成的 最长的回文串 。
// 在构造过程中，请注意 区分大小写 。比如 "Aa" 不能当做一个回文字符串。

// longestPalindrome .
// 出现偶数次的字符都可以作为回文串的字符，出现奇数次的数字可以选择其中的最多的偶数个作为回文串字符，且还可以再选择一个作为回文串中心字符
// 所以如果一个字符出现偶数次，则直接计入总结果，若出现奇数次，则取最多的偶数个计入总结果，并可以拿一个作为中心字符（中心字符全局只能有一个）
func longestPalindrome(s string) int {
	// 使用map统计各字符出现的次数
	hash := make(map[byte]int)
	for _, c := range []byte(s) {
		hash[c]++
	}

	// res用来记录偶数字符总个数，odd用来记录选择的中心字符（只能为0或1，因为中心字符只能有1个，只要有出现奇数次的字符，就将它置为1）
	res, odd := 0, 0
	for _, count := range hash {
		// 若当前字符出现奇数次，则将当前字符向下取偶数（这是可以出现在最终回文串中的），并将odd置为1表示最终的回文串有1个中间字符
		if count%2 != 0 {
			count--
			odd = 1
		}
		// 将偶数个字符数量加到总结果
		res += count
	}

	// 返回最终选择的偶数字符数量 + 中心字符（若有为1没有为0）
	return res + odd
}
