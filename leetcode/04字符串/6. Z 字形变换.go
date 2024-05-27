package main

import "strings"

// 6. Z 字形变换

// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

// convert .
// 行的遍历总是从0到numRows，再从numRows到0如此循环，因此遍历字符串s，将每个字符添加到其所在行中，
// 期间使用一个变量flag进行行的遍历方向控制，当flag为1时，字符添加方向为从0到numRows，flag为-1时，字符添加方向为从numRows到0
func convert(s string, numRows int) string {
	// 若只有一行，则直接返回s即可
	if numRows < 2 {
		return s
	}

	// 用来单独保存每一行的字符，在遍历结束后进行拼接
	rows := make([]*strings.Builder, numRows)
	for i := range rows {
		rows[i] = &strings.Builder{}
	}

	// 遍历每一个字符，将其添加到其所在行中，使用flag控制行的遍历方向
	i, flag := 0, -1 // i: 行号（当前字符应该添加到第几行）、flag: 控制行的遍历方向
	for _, c := range []byte(s) {
		// 将当前字符添加到其所在行中
		rows[i].WriteByte(c)

		// 打印到头调转方向
		if i == 0 || i == numRows-1 {
			flag *= -1
		}
		i += flag
	}

	// 将所有行的字符串进行拼接
	result := &strings.Builder{}
	for _, row := range rows {
		result.WriteString(row.String())
	}

	return result.String()
}
