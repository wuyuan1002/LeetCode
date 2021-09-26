package main

// 371. 两整数之和

// 给你两个整数 a 和 b ，不使用 运算符 + 和 - ，计算并返回两整数之和。

func main() {

}

// 位运算
// 两个二进制相加的所有可能: 0+0=0，0+1=1，1+0=1，1+1=0（进位） -- 相当于异或
func getSum1(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1 // 当前两个数字相加后进位的数字
		a ^= b                // 当前两个数字相加后不进位的数字
		b = carry
	}
	return a
}
