package main

import (
	"strconv"
)

// 165. 比较版本号

// 给你两个版本号 version1 和 version2 ，请你比较它们。
//
// 版本号由一个或多个修订号组成，各修订号由一个 '.' 连接。每个修订号由 多位数字 组成，
// 可能包含 前导零 。每个版本号至少包含一个字符。修订号从左到右编号，下标从 0 开始，
// 最左边的修订号下标为 0 ，下一个修订号下标为 1 ，以此类推。例如，2.5.33 和 0.1 都是有效的版本号。
//
// 比较版本号时，请按从左到右的顺序依次比较它们的修订号。比较修订号时，只需比较 忽略任何前导零后的整数值 。
// 也就是说，修订号 1 和修订号 001 相等 。如果版本号没有指定某个下标处的修订号，则该修订号视为 0 。
// 例如，版本 1.0 小于版本 1.1 ，因为它们下标为 0 的修订号相同，而下标为 1 的修订号分别为 0 和 1 ，0 < 1 。
//
// 返回规则如下：
// 如果version1>version2返回1，
// 如果version1<version2 返回 -1，
// 除此之外返回 0。

// func main() {

// }

// 1. 将两个版本号按照.分割成子串，分别从左到右比较子串
// 2. 各定义一个指针，返回到下一个.的子串并比较
func compareVersion(version1 string, version2 string) int {
	p, q := 0, 0             // 两个指针，分别指向两个版本号
	split1, split2 := "", "" // 两个版本号的下一个修订号
	done := false            // 是否已遍历到结束
	getNextSplit := func() { // 获取下一个修订号
		if p >= len(version1) && q >= len(version2) {
			done = true
			return
		}

		if p >= len(version1) {
			split1 = "0"
		} else {
			lastP := p
			for ; p < len(version1) && version1[p] != '.'; p++ {
			}
			split1 = version1[lastP:p]
			p++
		}

		if q >= len(version2) {
			split2 = "0"
		} else {
			lastQ := q
			for ; q < len(version2) && version2[q] != '.'; q++ {
			}
			split2 = version2[lastQ:q]
			q++
		}
	}

	for !done { // 比较当前修订号
		getNextSplit()
		v1, _ := strconv.Atoi(split1)
		v2, _ := strconv.Atoi(split2)

		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}

	return 0
}
