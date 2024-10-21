package main

import "strings"

// 71. 简化路径

// 给你一个字符串 path ，表示指向某一文件或目录的 Unix 风格 绝对路径 （以 '/' 开头），请你将其转化为更加简洁的规范路径。
// 在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..） 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。任意多个连续的斜杠（即，'//'）都被视为单个斜杠 '/' 。 对于此问题，任何其他格式的点（例如，'...'）均被视为文件/目录名称。

// simplifyPath .
//
// 将给定路径按照"/"分割，然后将各个目录压入栈中，遇到".."则弹出栈顶，最后返回栈中元素
func simplifyPath(path string) string {
	stack := make([]string, 0)
	// 遍历路径中的各个目录
	for _, name := range strings.Split(path, "/") {
		if name == ".." {
			// 若当前目录是".."且栈内目录数大于0，表示切换到父目录 -- 弹出栈顶目录
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if name != "" && name != "." {
			// 若当前目录不为空且不为"."表示切换至当前目录 -- 将当前目录入栈
			stack = append(stack, name)
		}
	}

	// 使用"/"拼接，返回栈内目录路径
	return "/" + strings.Join(stack, "/")
}
