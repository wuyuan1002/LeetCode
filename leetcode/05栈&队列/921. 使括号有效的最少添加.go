package main

// 921. 使括号有效的最少添加

// 只有满足下面几点之一，括号字符串才是有效的：
//
// 它是一个空字符串，或者
// 它可以被写成 AB （A 与 B 连接）, 其中 A 和 B 都是有效字符串，或者
// 它可以被写作 (A)，其中 A 是有效字符串。
// 给定一个括号字符串 s ，在每一次操作中，你都可以在字符串的任何位置插入一个括号
//
// 例如，如果 s = "()))" ，你可以插入一个开始括号为 "(()))" 或结束括号为 "())))" 。
// 返回 为使结果字符串 s 有效而必须添加的最少括号数。

// minAddToMakeValid .
// 遇到左括号则入栈，遇到右括号则出栈
// - 若遇到右括号时栈为空，则说明该右括号前面没有与其对应的左括号，所以需要为其添加一个左括号使得括号字符串有效
// - 若遍历结束后栈不为空，说明这些是没有与其匹配的右括号，则需要为栈内剩余的每个左括号添加一个右括号使得括号字符串有效
func minAddToMakeValid(s string) int {
	missLeft := 0            // 用来统计有多少缺失匹配左括号的右括号
	stack := make([]byte, 0) // 使用栈保存遍历过程中的左括号

	for _, c := range []byte(s) {
		if c == '(' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				// 当前右括号没有与其匹配的左括号
				missLeft++
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	// 最终栈内剩余的左括号是缺失匹配右括号的左括号数量 -- 总结果 = 缺失匹配右括号的左括号数量 + 缺失匹配左括号的右括号数量
	return len(stack) + missLeft
}

// minAddToMakeValid1 .
// 因为本题的栈只是用来表示当前右括号前面是否有多余的左括号与其匹配，
// 所以可以使用一个score用来对出现的左括号进行计数而不是使用栈真实存储遇到的左括号，
// 可以节省空间开销
func minAddToMakeValid1(s string) int {
	missLeft := 0  // 用来统计有多少缺失匹配左括号的右括号
	extraLeft := 0 // 使用一个变量用来记录当前有多少还未匹配右括号的左括号

	for _, c := range []byte(s) {
		if c == '(' {
			extraLeft++
		} else {
			if extraLeft == 0 {
				missLeft++
			} else {
				extraLeft--
			}
		}
	}

	return extraLeft + missLeft
}
