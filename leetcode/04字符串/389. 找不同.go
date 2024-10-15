package main

// 389. 找不同

// 给定两个字符串 s 和 t ，它们只包含小写字母。
// 字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
//
// 请找出在 t 中被添加的字母。

// findTheDifference .
//
// 将s和t的所有字符ascii码求和，两个和的差值就是添加字符的ascii码
func findTheDifference(s string, t string) byte {
	sum := 0
	for _, ch := range t {
		sum += int(ch)
	}
	for _, ch := range s {
		sum -= int(ch)
	}
	return byte(sum)
}
