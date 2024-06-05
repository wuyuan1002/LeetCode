package main

// LCR 169. 招式拆解 II

// 某套连招动作记作仅由小写字母组成的序列 arr，其中 arr[i] 第 i 个招式的名字。
// 请返回第一个只出现一次的招式名称，如不存在请返回空格。

// dismantlingAction169 .
// 返回第一个只出现一次的字符
func dismantlingAction169(arr string) byte {
	hmap := make(map[byte]int, len(arr))

	// 第一次遍历字符统计出现次数
	for _, c := range []byte(arr) {
		hmap[c]++
	}
	// 第二次遍历字符返回第一个只出现一次的字符
	for _, c := range []byte(arr) {
		if hmap[c] == 1 {
			return c
		}
	}

	return ' '
}
