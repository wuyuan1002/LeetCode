package main

import "strings"

// 796. 旋转字符串

// 给定两个字符串, s 和 goal。如果在若干次旋转操作之后，s 能变成 goal ，那么返回 true 。
//
// s 的 旋转操作 就是将 s 最左边的字符移动到最右边。
// 例如, 若 s = 'abcde'，在旋转一次之后结果就是'bcdea' 。

// rotateString .
// 字符串旋转即将字符串前一部分切割后拼接到原字符串末尾，
// 所以对于字符串s，将两个s进行相加后的s+s就包含了所有s可以通过旋转操作得到的字符串,
// 此时只需要判断goal是否在s+s中即可
func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}
