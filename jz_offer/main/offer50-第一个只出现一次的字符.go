package main

// 剑指 Offer 50. 第一个只出现一次的字符

// 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
//
// 示例:
// s = "abaccdeff"
// 返回 "b"
//
// s = ""
// 返回 " "

func main() {
}

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}

	num := make(map[byte]int32, 0) // 保存每一个字符出现的次鼠

	sbyte := []byte(s)
	for _, c := range sbyte {
		sum := num[c]
		num[c] = sum + 1
	}
	for _, c := range sbyte {
		if num[c] == 1 {
			return c
		}
	}
	return ' '
}

// 第二次做 -- 字符和次数存入map
func firstUniqChar1(s string) byte {
	if s == "" {
		return ' '
	}
	nums := make(map[byte]int, 0)
	sbyte := []byte(s)
	for _, c := range sbyte { // 第一次遍历，将字符和次数存入map
		nums[c] += +1
	}
	for _, c := range sbyte { // 第二次遍历，找出次数为1的字符
		if nums[c] == 1 {
			return c
		}
	}

	return ' '
}
