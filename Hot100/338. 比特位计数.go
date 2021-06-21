package main

// 338. 比特位计数

// 给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，
// 计算其二进制数中的 1 的数目并将它们作为数组返回。

func main() {

}

// 同offer 15
// 计算二进制中1的个数
// 1. 使用1和数字做与运算，之后每次把1左移一位
// 2. 每个数和它-1的数字做与运算，可以把该整数的最右面的1变成0，那么该整数有多少个1就可以做多少次这样的运算
func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ { // 遍历每一个数字
		num := i
		count := 0
		for num != 0 { // 计算当前数字中1的个数
			count++
			num = num & (num - 1)
		}
		result[i] = count
	}
	return result
}
