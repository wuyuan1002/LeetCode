package main

// 678. 有效的括号字符串

// 给定一个只包含三种字符的字符串：（，）和 *，写一个函数来检验这个字符串是否为有效字符串。有效字符串具有如下规则：
//
// 任何左括号 (必须有相应的右括号 )。
// 任何右括号 )必须有相应的左括号 (。
// 左括号 ( 必须在对应的右括号之前 )。
// *可以被视为单个右括号 )，或单个左括号 (，或一个空字符串。
// 一个空字符串也被视为有效字符串。

func main() {

}

// 类似20
// 定义两个栈分别存左括号和*
// 1. 如果遇到左括号，则将当前下标存入左括号栈
// 2. 如果遇到星号，则将当前下标存入星号栈
// 3. 如果遇到右括号，则需要有一个左括号或星号和右括号匹配，由于星号也可以看成右括号或者空字符串，因此当前的右括号应优先和左括号匹配，没有左括号时和星号匹配
func checkValidString(s string) bool {
	if s == "" {
		return false
	}

	lstack, xstack := make([]int, 0), make([]int, 0)
	for i, c := range s {
		if c == '(' {
			lstack = append(lstack, i)
		} else if c == '*' {
			xstack = append(xstack, i)
		} else if c == ')' {
			if len(lstack) > 0 {
				// 优先匹配左括号
				lstack = lstack[:len(lstack)-1]
			} else if len(xstack) > 0 {
				// 若左括号没有了则使用*号匹配
				xstack = xstack[:len(xstack)-1]
			} else {
				// 若都没有则说明没有左括号与当前右括号匹配
				return false
			}
		}
	}

	// 从后向前检查每一个左括号和与其匹配的*号
	// 此时将*号视为右括号或空串，且每个左括号必须出现在与其匹配的*号之前
	i := len(lstack) - 1
	for j := len(xstack) - 1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if lstack[i] > xstack[j] {
			return false
		}
	}

	return i < 0
}
