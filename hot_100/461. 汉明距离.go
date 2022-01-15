package main

// 461. 汉明距离

// 两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
//
// 给你两个整数 x 和 y，计算并返回它们之间的汉明距离。

// func main() {

// }

// 类似offer 15，Hot100 338
// 先让两个数异或，得到为1的位置就是两数不同的位，之后统计1的个数
func hammingDistance(x int, y int) int {
	num := x ^ y // 两数求异或
	count := 0
	for num != 0 { // 计算异或后得数的1的个数
		num &= num - 1
		count++
	}
	return count
}
