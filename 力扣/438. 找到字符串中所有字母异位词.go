package main

// 438. 找到字符串中所有字母异位词

// func main() {

// }

// 滑动窗口+(数组或map)
func findAnagrams(s string, p string) []int {
	if s == "" || p == "" || len(s) < len(p) {
		return nil
	}

	res := make([]int, 0)
	pcnt := make([]byte, 26)
	scnt := make([]byte, 26)

	isEqual := func() bool {
		for i := 0; i < 26; i++ {
			if scnt[i] != pcnt[i] {
				return false
			}
		}
		return true
	}

	for _, c := range p {
		pcnt[c-'a']++
	}

	for i := 0; i <= len(s); i++ {
		if i < len(p) {
			scnt[s[i]-'a']++
			continue
		}

		if isEqual() {
			res = append(res, i-len(p))
		}

		if i < len(s) {
			scnt[s[i]-'a']++
			scnt[s[i-len(p)]-'a']--
		}
	}

	return res
}
