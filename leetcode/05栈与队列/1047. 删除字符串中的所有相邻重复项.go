package main

// 1047. 删除字符串中的所有相邻重复项

// 给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
// 在 S 上反复执行重复项删除操作，直到无法继续删除。
//
// 在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

// removeDuplicates .
// 使用栈保存遍历过程中的字符, 遍历下一个字符时与栈顶元素比较, 若相同则出栈, 最终栈内元素就是删除后的结果
func removeDuplicates(s string) string {
	if len(s) < 2 {
		return s
	}

	stack := make([]byte, 0)
	for _, c := range []byte(s) {
		if len(stack) > 0 && c == stack[len(stack)-1] {
			// 若与栈顶元素相等则出栈继续遍历下一个字符
			stack = stack[:len(stack)-1]
		} else {
			// 若不想等则将当前元素入栈
			stack = append(stack, c)
		}
	}

	// 最终栈内元素即为删除后的结果
	return string(stack)
}
