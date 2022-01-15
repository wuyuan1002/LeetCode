package main

// 5. 最长回文子串

// func main() {
// 	fmt.Println(longestPalindrome("babad"))
// }

// 中心扩散法
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	longest := "" // 最长回文子串
	for i := 0; i < len(s); i++ {
		l, r := i-1, i+1
		for l >= 0 && s[l] == s[i] {
			l--
		}
		for r < len(s) && s[r] == s[i] {
			r++
		}
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > len(longest) {
			longest = s[l+1 : r]
		}
	}
	return longest
}
