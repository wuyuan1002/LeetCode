package main

import (
	"fmt"
)

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
// 注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。

func main() {
	fmt.Println(minWindow("abbbcdbbefg", "bf"))
}

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	sbyte := []byte(s)
	tbyte := []byte(t)

	need := make(map[byte]int)
	for _, v := range tbyte {
		need[v]++
	}

	isAllExist := func() bool { // 校验子串是否全部包含了map中的值
		for _, v := range need {
			if v > 0 {
				return false
			}
		}
		return true
	}

	minl, minr := 0, 0 // 某一次寻找时的结果
	resl, resr := 0, 0 // 最终结果
	finded := false    // 是否存在子串
	l, r := 0, 0       // 遍历时的左右指针
	for ; r < len(sbyte); r++ {
		if _, ok := need[sbyte[r]]; !ok {
			continue
		}

		need[sbyte[r]]--
		for ; isAllExist() && r-l+1 >= len(tbyte); l++ {
			// 标记起码找到了一个子串
			finded = true

			// 记录当前r的位置
			minr = r
			if _, ok := need[sbyte[l]]; ok && need[sbyte[l]] <= 0 {
				// 若l所在位置的字符可以被舍弃
				need[sbyte[l]]++
				minl = l
			}
		}

		if resr-resl > minr-minl { // 若查询到的新子串比原来的子串更小，更新结果
			resl = minl
			resr = minr
		} else if resl == 0 && resr == 0 { // resl和resr还是初始值的时候，直接赋值
			resl = minl
			resr = minr
		}
	}

	if !finded {
		return ""
	}
	return string(sbyte[resl : resr+1])
}
