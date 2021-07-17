package main

import (
	"fmt"
	"strconv"
)

// 394. 字符串解码

// 给定一个经过编码的字符串，返回它解码后的字符串。
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像3a或2[4]的输入。

func main() {
	fmt.Println(decodeString("as13[a]2[bc]zx"))
}

// 使用栈，遇到左括号入栈，遇到右括号出栈
func decodeString(s string) string {
	if s == "" {
		return s
	}

	// 栈内元素
	type entry struct {
		num int
		str string
	}

	stack := []*entry{{
		num: 1,
		str: "",
	}}

	num := ""
	for _, c := range s {
		if c >= '0' && c <= '9' {
			num += string(c)
		} else if c == '[' {
			n, _ := strconv.Atoi(num)
			stack = append(stack, &entry{
				num: n,
			})
			num = ""
		} else if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			stack[len(stack)-1].str += string(c)
		} else if c == ']' {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			str := top.str
			for i := 1; i < top.num; i++ {
				top.str += str
			}
			stack[len(stack)-1].str += top.str
		} else {
			panic("遇到非法字符")
		}
	}

	return stack[len(stack)-1].str
}
