package main

import (
	"fmt"
	"strconv"
)

// 394. 字符串解码

// 给定一个经过编码的字符串，返回它解码后的字符串。
//
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。
// 注意 k 保证为正整数。
//
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像3a或2[4]的输入。

func main() {
	// fmt.Println(decodeString("as3[a]2[bc]zx"))
	// fmt.Println(decodeString("dhj3[a2[c]]"))
	fmt.Println(decodeString("12[ads]"))
}

// 使用栈，遇到左括号入栈，遇到右括号出栈
func decodeString(s string) string {
	if s == "" {
		return ""
	}

	// 栈内元素
	type entry struct {
		num int
		str string
	}

	// 定义栈
	stack := make([]*entry, 0)
	stack = append(stack, &entry{ // 先在栈内放一个元素，存最终的结果
		num: 1,
		str: "",
	})

	num := ""
	for _, c := range s {
		if c >= '0' && c <= '9' {
			// 若是数字 -- 计算下一个左括号前的数字大小 -- 如ab100[asd]
			num += string(c)
		} else if c == '[' {
			// 若是左括号 -- 入栈一个新元素
			n, _ := strconv.Atoi(num)
			num = ""
			stack = append(stack, &entry{
				num: n,
				str: "",
			})
		} else if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			// 若是字母 -- 将字母加到栈顶元素上
			stack[len(stack)-1].str += string(c)
		} else if c == ']' {
			// 若是右括号 -- 弹出一个栈顶元素，计算它的字符串，并把结果加到新的栈顶元素上
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			str := top.str
			for i := 0; i < top.num-1; i++ {
				top.str += str
			}

			stack[len(stack)-1].str += top.str
		}
	}

	// 最终栈内会只剩下刚开始栈内加入的那个用来保存最终结果的元素
	return stack[len(stack)-1].str
}
