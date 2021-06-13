package main

import (
	"fmt"
)

// 17. 电话号码的字母组合

// 给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

func main() {
	fmt.Println(letterCombinations("23"))
}

// 定义数字字母对应表
var phone = map[string][]byte{
	"2": {'a', 'b', 'c'},
	"3": {'d', 'e', 'f'},
	"4": {'g', 'h', 'i'},
	"5": {'j', 'k', 'l'},
	"6": {'m', 'n', 'o'},
	"7": {'p', 'q', 'r', 's'},
	"8": {'t', 'u', 'v'},
	"9": {'w', 'x', 'y', 'z'},
}

// offer 17
// 回溯法
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	result := make([]string, 0) // 所有结果
	res := make([]byte, 0)      // 某一次遍历的结果的结果
	dfs(digits, 0, &res, &result)
	return result
}

func dfs(digits string, index int, res *[]byte, result *[]string) {

	// 如果已经到最后一个数字，则将当前结果存入总结果中
	if index == len(digits) {
		*result = append(*result, string(*res))
		return
	}

	// 获取当前数字
	number := string(digits[index])

	// 遍历当前数字对应的所有字符
	for _, char := range phone[number] {
		// 先把当前字符放入结果中
		*res = append(*res, char)
		// 遍历下一个数字的所有字符
		dfs(digits, index+1, res, result)
		// 将当前字符从结果中去除
		*res = (*res)[:len(*res)-1]
	}
}
