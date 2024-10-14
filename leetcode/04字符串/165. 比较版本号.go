package main

import (
	"strconv"
	"strings"
)

// 165. 比较版本号

// 给你两个 版本号字符串 version1 和 version2 ，请你比较它们。版本号由被点 '.' 分开的修订号组成。
// 修订号的值 是它 转换为整数 并忽略前导零。
//
// 比较版本号时，请按 从左到右的顺序 依次比较它们的修订号。
// 如果其中一个版本字符串的修订号较少，则将缺失的修订号视为 0。
//
// 返回规则如下：
//
// 如果 version1 < version2 返回 -1，
// 如果 version1 > version2 返回 1，
// 除此之外返回 0。

// compareVersion .
//
// 将字符串按照"."进行分割，然后从左至右挨个进行数字比较，
// 若发现一个版本的修订号少则按0处理（缺失的修订号在从右侧进行补0）
func compareVersion(version1, version2 string) int {
	// 将两个版本字符串按照"."进行分割
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	// 从左到右挨个比较对应版本号的大小，缺失的进行补0操作
	for i := 0; i < len(v1) || i < len(v2); i++ {
		x, y := 0, 0
		if i < len(v1) {
			x, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			y, _ = strconv.Atoi(v2[i])
		}
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}
