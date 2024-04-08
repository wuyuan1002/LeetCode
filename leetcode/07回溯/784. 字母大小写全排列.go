package main

// 784. 字母大小写全排列

// 给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。
//
// 返回 所有可能得到的字符串集合 。以 任意顺序 返回输出。

// letterCasePermutation .
// 字母大小写转换规则，例如 char c：
// 	  互相转(大转小，小转大)，异或32：c ^ 32
// 	  大转小，小不变，或 32：c | 32
//    小转大，大不变，与非32：c & ~32
func letterCasePermutation(s string) []string {
	result := make([]string, 0) // 总结果集
	dfsLetterCasePermutation([]byte(s), 0, &result)
	return result
}

// dfsLetterCasePermutation
// 可以使用res来存储已经选择好的字符，也可以使用直接在原数组中转换字符的方式实现字符选择
//
// sb: 选择列表
// start: 每次遍历的起始下标 -- 指定当前层的选择范围 -- sb[start: ]
// result: 总结果集
func dfsLetterCasePermutation(sb []byte, start int, result *[]string) {
	// 每次都在选择列表中原地进行转换，所以每次回溯都做了一次转换，是一个有效的结果，先将本次结果加入到结果集
	*result = append(*result, string(sb))

	// 不断将选择列表中的字母进行大小写转换然后进入下一层回溯
	for i := start; i < len(sb); i++ {
		// 若当前字符是数字，则直接跳过
		if isDigit(sb[i]) {
			continue
		}

		// 若当前字符是字母，则进行大小写转换并进入下一层回溯，回溯结束后再将其转换回来
		sb[i] ^= 32
		dfsLetterCasePermutation(sb, i+1, result)
		sb[i] ^= 32
	}
}

// isDigit 判断给定的字符是否为数字
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
