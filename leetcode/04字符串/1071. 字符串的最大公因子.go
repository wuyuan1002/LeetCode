package main

// 1071. 字符串的最大公因子

// 对于字符串 s 和 t，只有在 s = t + t + t + ... + t + t（t 自身连接 1 次或多次）时，我们才认定 “t 能除尽 s”。
//
// 给定两个字符串 str1 和 str2 。返回 最长字符串 x，要求满足 x 能除尽 str1 且 x 能除尽 str2 。

// gcdOfStrings .
// 假设两个字符串存在最大公因子，该最大公因子为X，长度为N，说明str1和str2都由特定数量的X构成，
// 即 str1 = p * X, len(str1) = p * N 、 str2 = q * X, len(str2) = q * N
// 说明两字符串的长度一定为N的整数倍，即N为两字符串长度的最大公因数，
// 且 p*X + q*X == q*X + p*X，即 str1+str2 == str2+str1
func gcdOfStrings(str1 string, str2 string) string {
	// 若两字符串相加后不相等，说明不存在最大公因子
	if str1+str2 != str2+str1 {
		return ""
	}

	// 存在最大公因子，且最大公因子长度为两字符串长度的最大公因数
	// 那么一个字符串的前最大公因数个字符的子串就是两字符串的字符串的最大公因子
	return str1[:gcd(len(str1), len(str2))]
}

// gcd 辗转相除法求最大公约数
func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}
