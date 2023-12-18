package main

// 17. 电话号码的字母组合

// 给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

// letterCombinations .
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	result := make([]string, 0) // 总结果
	res := make([]byte, 0)      // 存放一次回溯路径的结果 -- 一次回溯的字母组合

	dfsLetterCombinations(digits, 0, &res, &result)
	return result
}

// dfsLetterCombinations .
// digits：选择列表 -- 数字列表
// index：指定当前层的数字下标 -- 当前层的字母选择范围即为phone[string(digits[index])]
// res: 存放一次回溯路径的结果 -- 一次回溯的字母组合
// result: 总结果 -- 所有的字母组合
func dfsLetterCombinations(digits string, index int, res *[]byte, result *[]string) {
	// 已经遍历到最后一个数字，说明已经找到了一个字母组合，将回溯路径的结果记录到结果集
	if index == len(digits) {
		*result = append(*result, string(*res))
		return
	}

	// 获取当前层的数字
	number := digits[index]

	// 遍历当前数字对应的所有字符
	for _, char := range phone[number] {
		// 先把当前字符放入回溯路径中
		*res = append(*res, char)
		// 递归进入下一层 -- 遍历下一个数字的所有字符
		dfsLetterCombinations(digits, index+1, res, result)
		// 将当前字符从回溯路径中去除
		*res = (*res)[:len(*res)-1]
	}
}

// phone 数字与字母对应表
var phone = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}
