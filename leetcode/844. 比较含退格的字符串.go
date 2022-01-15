package main

// 844. 比较含退格的字符串

// 给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，
// 判断二者是否相等，并返回结果。 # 代表退格字符。
//
// 注意：如果对空文本输入退格字符，文本继续为空。

// func main() {

// }

// 1. 使用栈存储字符，遇到字符入栈，遇到#弹出栈顶元素，最后比较栈内元素是否相等
// 2. 双指针
func backspaceCompare(s, t string) bool {
	return build(s) == build(t)
}

func build(str string) string {
	s := make([]byte, 0)
	for i := range str {
		if str[i] != '#' {
			s = append(s, str[i])
		} else if len(s) > 0 {
			s = s[:len(s)-1]
		}
	}
	return string(s)
}
