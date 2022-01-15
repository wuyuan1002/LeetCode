package main

// 76. 最小覆盖子串

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
// 注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。

// func main() {
// 	// fmt.Println(minWindow("bba", "ab"))
// 	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
// }

// 超出时间限制
func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}

	need := make(map[byte]int, len(t))
	for _, c := range t {
		need[byte(c)]++
	}

	isAllExist := func() bool {
		for _, v := range need {
			if v > 0 {
				return false
			}
		}
		return true
	}

	subStr := "" // 子串
	l, r := 0, 0 // 左右指针
	for ; r < len(s); r++ {
		if _, ok := need[s[r]]; !ok {
			continue
		}

		need[s[r]]--
		for ; isAllExist() && r-l+1 >= len(t); l++ {
			if _, ok := need[s[l]]; ok && need[s[l]] <= 0 {
				if subStr == "" || len(subStr) > r-l+1 {
					subStr = s[l : r+1]
				}
				need[s[l]]++
			}
		}
	}

	return subStr
}
