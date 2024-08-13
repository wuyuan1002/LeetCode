package main

// LCR 167. 招式拆解 I

// 某套连招动作记作序列 arr，其中 arr[i] 为第 i 个招式的名字。
// 请返回 arr 中最多可以出连续不重复的多少个招式。

// dismantlingAction167 .
// leetcode 3. 无重复字符的最长子串
//
// 双指针滑动窗口，同时使用map存储字符最近一次出现的下标
// 滑动窗口的的字符所构成的子串就是每一个无重复字符的子串，求无重复字符的最长子串长度
func dismantlingAction167(s string) int {
	sbyte := []byte(s)
	hash := make(map[byte]int) // 存储字符和它最近一次出现的下标
	maxLen := 0                // 无重复字符的最长子串长度

	// 双指针遍历字符串形成滑动窗口，同时根据窗口长度计算无重复字符的最长子串长度
	for l, r := 0, 0; r < len(sbyte); r++ {
		// 若当前字符已经出现过，将左边界向右移动或保持不变，不能直接移动到index+1，原因: abba
		// 当前字符上次出现的位置可能在窗口内也可能不在窗口内，若在窗口内那么需要向右移动左指针到index+1，若没在窗口内那就不需要移动左指针
		if index, ok := hash[sbyte[r]]; ok {
			l = max(l, index+1)
		}

		// 记录或更新当前字符最近一次出现的下标
		hash[sbyte[r]] = r

		// 更新无重复字符的最长子串长度
		maxLen = max(maxLen, r-l+1)
	}

	return maxLen
}
