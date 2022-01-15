package main

import "strings"

// 1249. 移除无效的括号

// 给你一个由 '('、')' 和小写字母组成的字符串 s。
//
// 你需要从字符串中删除最少数目的 '(' 或者 ')'（可以删除任意位置的括号)，使得剩下的「括号字符串」有效。
//
// 请返回任意一个合法字符串。
//
// 有效「括号字符串」应当符合以下任意一条要求：
//
// 空字符串或只包含小写字母的字符串
// 可以被写作AB（A连接B）的字符串，其中A和B都是有效「括号字符串」
// 可以被写作(A)的字符串，其中A是一个有效的「括号字符串」

// func main() {

// }

// 使用栈来保存遇到的左括号出现的下标，每次遇到右括号就从栈顶出栈
func minRemoveToMakeValid(s string) string {
	if s == "" {
		return s
	}

	stack := make([]int, 0)             // 使用栈来保存遇到的左括号出现的下标
	remove := make(map[int]struct{}, 0) // 保存待删除的括号的下标
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else if c == ')' {
			if len(stack) > 0 { // 若有左括号，则抵消掉一个左括号
				stack = stack[:len(stack)-1]
			} else { // 若没有左括号，则记录下来需要被删除
				remove[i] = struct{}{}
			}
		}
	}

	// 多余的左括号也需要被删除
	for _, i := range stack {
		remove[i] = struct{}{}
	}

	// 重新构建字符串，删除掉多余的括号
	sb := strings.Builder{}
	for i, c := range s {
		if _, ok := remove[i]; !ok {
			sb.WriteByte(byte(c))
		}
	}
	return sb.String()
}
