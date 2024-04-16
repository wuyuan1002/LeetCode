package main

// 1021. 删除最外层的括号

// 有效括号字符串为空 ""、"(" + A + ")" 或 A + B ，其中 A 和 B 都是有效的括号字符串，+ 代表字符串的连接。
//
// 例如，""，"()"，"(())()" 和 "(()(()))" 都是有效的括号字符串。
// 如果有效字符串 s 非空，且不存在将其拆分为 s = A + B 的方法，我们称其为原语（primitive），其中 A 和 B 都是非空有效括号字符串。
//
// 给出一个非空有效字符串 s，考虑将其进行原语化分解，使得：s = P_1 + P_2 + ... + P_k，其中 P_i 是有效括号字符串原语。
//
// 对 s 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 s 。

// removeOuterParentheses .
// 使用栈遇到左括号入栈遇到右括号出栈，当栈为空时说明找到了一个原语，
// 要将每个原语的外层括号去除，也就是该原语入栈的第一个左括号和最后一个入栈的右括号去除
func removeOuterParentheses(s string) string {
	result := make([]byte, 0)
	stack := make([]byte, 0)

	for _, c := range []byte(s) {
		// 因为第一个入栈的左括号和最后一个出栈的右括号不加入到结果集，
		// 所以这里一定是先判断是否为右括号再判断是否为左括号，中间根据栈中是否有元素判断是否要将当前括号加入到结果集，
		// 因为若是第一个左括号，此时栈为空，先由于栈为空使其无法加入到结果集然后再将其加入到栈中，
		// 若是最后一个右括号，此时栈中只有一个左括号了，先将其从栈中去除，然后再判断栈是否为空时由于为空使当前右括号也无法加入到结果集中，
		// 这样就实现了去除原语的第一个左括号和最后一个的右括号
		if c == ')' {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			result = append(result, c)
		}
		if c == '(' {
			stack = append(stack, c)
		}
	}

	return string(result)
}

// removeOuterParentheses1 .
// 优化：因为上面使用栈的解法，使用栈只是根据栈是否为空来判断当前左右括号是否为栈的最外层括号，
// 因此可以直接使用一个变量leftCount来表示当前原语中遇到左括号的数量，当其为0时说明当前原语遍历结束，
// leftCount为0表示栈为空
func removeOuterParentheses1(s string) string {
	result := make([]byte, 0)
	leftCount := 0

	for _, c := range []byte(s) {
		if c == ')' {
			leftCount--
		}
		if leftCount > 0 {
			result = append(result, c)
		}
		if c == '(' {
			leftCount++
		}
	}

	return string(result)
}
