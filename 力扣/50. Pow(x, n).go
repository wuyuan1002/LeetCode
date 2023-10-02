package main

// 50. Pow(x, n)

// 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。

// func main() {

// }

// 同 offer 16
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	isNeg := n < 0 // 次方是否为负数
	if isNeg {
		n = -n
	}

	res := myPow(x, n>>1)
	res *= res
	if n&0x1 == 1 {
		res *= x
	}

	if isNeg {
		res = 1 / res
	}
	return res
}
