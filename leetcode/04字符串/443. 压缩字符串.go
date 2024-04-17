package main

import "strconv"

// 443. 压缩字符串

// 给你一个字符数组 chars ，请使用下述算法压缩：
//
// 从一个空字符串 s 开始。对于 chars 中的每组 连续重复字符 ：
//
// 如果这一组长度为 1 ，则将字符追加到 s 中。
// 否则，需要向 s 追加字符，后跟这一组的长度。
// 压缩后得到的字符串 s 不应该直接返回 ，需要转储到字符数组 chars 中。需要注意的是，如果组长度为 10 或 10 以上，则在 chars 数组中会被拆分为多个字符。
//
// 请在 修改完输入数组后 ，返回该数组的新长度。
//
// 你必须设计并实现一个只使用常量额外空间的算法来解决此问题。

// compress .
// 双指针，左指针指向答案待插入的位置（[0, l)范围是已确定的结果），右指针指向当前处理的字符位置
func compress(chars []byte) int {

	l, r := 0, 0
	for r < len(chars) {
		// start表示当前字符的初始位置，不断向后移动右指针，直到其指向不等于当前字符的位置
		start := r
		for r < len(chars) && chars[r] == chars[start] {
			r++
		}

		// 统计当前字符连续出现的次数
		count := r - start

		// 将当前字符写入结果
		chars[l] = chars[start]
		l++

		// 将当前字符连续出现的次数写入结果 -- 只有当连续出现次数大于1时才在后面跟出现次数，只出现一次的不跟出现次数
		if count > 1 {
			for _, n := range toCharArray(count) {
				chars[l] = n
				l++
			}
		}
	}

	return l
}

// toCharArray 将给定数字每一位转成字符并存入byte数组后返回 -- 如 123 -> ['1', '2', '3']
func toCharArray(num int) []byte {
	numStr := strconv.Itoa(num)
	array := make([]byte, len(numStr))
	for i := range numStr {
		array[i] = numStr[i]
	}
	return array
}
