package main

// 231. 2的幂

// 给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。
//
// 如果存在一个整数 x 使得n == 2的x次方 ，则认为 n 是 2 的幂次方。

// func main() {

// }

// 类似 offer15 二进制中1的个数
// 若n为正整数且二进制位中只有一个1就说明n是2的幂次方
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
