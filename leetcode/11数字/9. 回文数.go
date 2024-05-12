package main

// 9. 回文数

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
// 回文数
// 是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
// 例如，121 是回文，而 123 不是。

// isPalindrome .
// 1. 转成字符串使用双指针进行判断
// 2. 将数字按照倒序重新获取一个新的数字，然后比较两数字是否相等（若原数字是负数则一定不是回文数，因为 "-121" -> "121-"）
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := 0
	for num := x; num != 0; num /= 10 {
		y = y*10 + num%10
	}

	return y == x
}
