package main

// 187. 重复的DNA序列

// 所有 DNA 都由一系列缩写为 'A'，'C'，'G' 和 'T' 的核苷酸组成，例如："ACGAATTCCG"。
// 在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。
//
// 编写一个函数来找出所有目标子串，目标子串的长度为 10，且在 DNA 字符串 s 中出现次数超过一次。

// func main() {
// 	findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
// }

// 使用map存所有长度为10的子串和出现的次数
func findRepeatedDnaSequences(s string) []string {
	mapp := make(map[string]int)
	res := make([]string, 0)
	for i := 0; i <= len(s)-10; i++ {
		subs := s[i : i+10]  // 子串
		mapp[subs]++         // 出现次数+1
		if mapp[subs] == 2 { // 若出现次数超过2次
			res = append(res, subs)
		}
	}
	return res
}
