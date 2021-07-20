package main

// 9. 回文数

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

func main() {

}

// 负数肯定不是回文数
// 1. 转成字符串判断是否为回文串
// 2. 将倒叙数字算出来，和原数字比较是否相等
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}

	cur := 0
	num := x
	for num != 0 {
		cur = cur*10 + num%10
		num /= 10
	}

	return cur == x
}
