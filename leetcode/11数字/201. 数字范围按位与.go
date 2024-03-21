package main

// 201. 数字范围按位与

// 给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）。

// rangeBitwiseAnd .
// 本题其实是求给定范围数字的二进制位的公共前缀
// 把一个整数减1再和原来的整数作与运算，会把该整数最右面的1变成0
// 对于[n, m]范围的数字，公共前缀所构成的数字，一定小于或等于n，因此对m进行以上操作，清除最右边的1，直到它小于或等于n，此时的m便是公共前缀所构成的数字，返回该数字即可
func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right &= right - 1
	}
	return right
}
