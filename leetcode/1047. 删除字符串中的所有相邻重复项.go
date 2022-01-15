package main

// 1047. 删除字符串中的所有相邻重复项

// 给出由小写字母组成的字符串S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
// 在 S 上反复执行重复项删除操作，直到无法继续删除。
//
// 在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

// func main() {

// }

// 使用栈
func removeDuplicates2(s string) string {
	if len(s) < 2 {
		return s
	}

	stack := make([]byte, 0)
	for _, c := range s {
		if len(stack) > 0 && byte(c) == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(c))
		}
	}
	return string(stack)
}
