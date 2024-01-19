package main

import "strconv"

// 394. 字符串解码

// 给定一个经过编码的字符串，返回它解码后的字符串。
//
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

// decodeString .
// 使用栈，遇到左括号入栈，遇到右括号出栈
func decodeString(s string) string {
	// 栈内元素 -- 用于表示str重复num次
	type entry struct {
		num int
		str string
	}

	// 定义栈
	stack := make([]*entry, 0)
	stack = append(stack, &entry{num: 1, str: ""}) // 先在栈内放一个元素，存最终的结果

	// 遍历字符串的每个字符
	num := ""
	for _, c := range s {
		if c >= '0' && c <= '9' {
			// 若是数字 -- 使用num记录下一个扩后内的数字出现几次
			num += string(c)
		} else if c == '[' {
			// 若是左括号 -- 入栈一个新元素，num表示该括号内的字符串循环几次
			n, _ := strconv.Atoi(num)
			stack = append(stack, &entry{num: n, str: ""})
			num = ""
		} else if c >= 'a' && c <= 'z' {
			// 若是字母 -- 表示为当前栈顶元素的字符串内容，将字母加到栈顶元素上
			stack[len(stack)-1].str += string(c)
		} else if c == ']' {
			// 若是右括号 -- 弹出一个栈顶元素，计算它的字符串，并把结果加到新的栈顶元素上
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 重复该括号内的字符串
			str := top.str
			for i := 0; i < top.num-1; i++ {
				top.str += str
			}

			// 将目标字符串加到新的栈顶元素的字符串上
			stack[len(stack)-1].str += top.str
		}
	}

	// 遍历结束后，栈内只剩下刚开始存的一个元素，该元素的字符串内容就是结果
	return stack[0].str
}
