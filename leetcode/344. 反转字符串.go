package main

// 344. 反转字符串

func main() {

}

// 双指针
func reverseString(s []byte) {
	if s == nil || len(s) < 2 {
		return
	}

	l, r := 0, len(s)
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
