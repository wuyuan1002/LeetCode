package main

// 剑指 Offer 64. 求1+2+…+n

// 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

// func main() {

// }

// 不让用循环 - 那就用递归
// 不让用if - 那就用&&
var res int

func sumNums(n int) int {
	res = 0
	sum(n)
	return res
}

func sum(n int) bool {
	res += n
	return 0 < n && sum(n-1)
}

// func sumNums(n int) int {
// 	res := 0
// 	var sum func(int) bool
// 	sum = func(n int) bool {
// 		res += n
// 		return n > 0 && sum(n-1)
// 	}
// 	sum(n)
// 	return res
// }
