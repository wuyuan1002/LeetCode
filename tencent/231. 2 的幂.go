package main

// 231. 2的幂

// 给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。
//
// 如果存在一个整数 x 使得n == 2的x次方 ，则认为 n 是 2 的幂次方。

func main() {

}

// 类似 offer15 二进制中1的个数
// 若n的二进制位中初最低位外只有一个1就说明n是2的幂次方
func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 || n&0x1 == 1 {
		// 若n的第一位为1，说明n为奇数，肯定不是2的幂次方
		return false
	}

	// 判断n的二进制中1的个数是否只有一个
	count := 0
	for n != 0 {
		count++
		n &= n - 1
	}

	return count == 1
}
