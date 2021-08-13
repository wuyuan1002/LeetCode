package main

import "fmt"

// 1297. 子串的最大出现次数

// 给你一个字符串s ，请你返回满足以下条件且出现次数最大的任意子串的出现次数：
//
// 子串中不同字母的数目必须小于等于 maxLetters 。
// 子串的长度必须大于等于minSize 且小于等于maxSize 。

func main() {
	fmt.Println(maxFreq("aababcaab", 2, 3, 4))
}

// 若长度为maxSize的子串可以，那么它的子串也必然可以，所以必然存在minSize的子串满足
// 所以只需要求minSize长度的子串就可以了
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	if s == "" || len(s) < maxLetters {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	hash := make(map[string]int, 0) // 使用map存储子串出现的次数
	// 寻找满足条件的所有子串
	for l, r := 0, minSize-1; r < len(s); l, r = l+1, r+1 {
		letterMap := make(map[byte]int, minSize) // 用来判断子串中不同字母的出现次数
		for i := l; i <= r; i++ {
			letterMap[s[i]]++
		}
		if len(letterMap) <= maxLetters {
			// 当前子串不同字母个数小于 maxLetters，符合条件
			hash[s[l:r+1]]++
		}
	}

	// 求出子串出现最多的次数
	res := 0
	for _, v := range hash {
		res = max(res, v)
	}
	return res
}
