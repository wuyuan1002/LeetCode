package main

import "strings"

// 12. 整数转罗马数字

// 七个不同的符号代表罗马数字，其值如下：
//
// 符号	值
// I	1
// V	5
// X	10
// L	50
// C	100
// D	500
// M	1000
// 罗马数字是通过添加从最高到最低的小数位值的转换而形成的。将小数位值转换为罗马数字有以下规则：
//
// 如果该值不是以 4 或 9 开头，请选择可以从输入中减去的最大值的符号，将该符号附加到结果，减去其值，然后将其余部分转换为罗马数字。
// 如果该值以 4 或 9 开头，使用 减法形式，表示从以下符号中减去一个符号，例如 4 是 5 (V) 减 1 (I): IV ，9 是 10 (X) 减 1 (I)：IX。仅使用以下减法形式：4 (IV)，9 (IX)，40 (XL)，90 (XC)，400 (CD) 和 900 (CM)。
// 只有 10 的次方（I, X, C, M）最多可以连续附加 3 次以代表 10 的倍数。你不能多次附加 5 (V)，50 (L) 或 500 (D)。如果需要将符号附加4次，请使用 减法形式。
// 给定一个整数，将其转换为罗马数字。

// intToRoman .
//
// 罗马数字的基础组成只包括baseRoman中的数字，其它数字都是通过加法形式将它们进行拼接得到
//
// 从大到小不断遍历baseRoman，将baseRoman中的数字不断从num中减去，
// 然后将被减去数字对应的罗马数字的字符串拼接到总结果的后面，直到将num减到0后返回总结果字符串
func intToRoman(num int) string {
	result := &strings.Builder{}
	// 不断从大到小将基础数字从目标数字中减去并将罗马数字拼接到结果中
	for _, base := range baseRoman {
		for num >= base.num {
			num -= base.num
			result.WriteString(base.roman)
		}
		if num == 0 {
			break
		}
	}
	return result.String()
}

// baseRoman 罗马数字中的基本数字，其它数字都是由这里的数字相加得到
var baseRoman = []struct {
	num   int
	roman string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}
