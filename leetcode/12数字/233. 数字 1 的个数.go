package main

// 233. 数字 1 的个数

// 给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。

// countDigitOne .
// 将所有数字按照个位、十位、百位...分别进行个数的统计，并将总结果相加即可 -- 将1~n的个位、十位、百位...的1出现次数相加，即为1出现的总次数
// 将currentrent记为当前位，high记为currentrent的高位、low记为低位、digit记为位因子（如数字为3402，currentrent为0的话，digit为10、high为34、low为2）
// - 当current=0时： 此位1的出现次数只由高位high决定，公式为：high × digit
// - 当current=1时： 此位1的出现次数由高位high和低位low决定，公式为：high × digit + low + 1
// - 当current=2,3,⋯,9时： 此位1的出现次数只由高位high决定，公式为：(high + 1) × digit
func countDigitOne(n int) int {
	// 总结果 -- 每一位上1出现的次数和
	result := 0

	// 开始计算每一位，按照个位、十位、百位...进行计算
	for digit, current, high, low := 1, n%10, n/10, 0; high != 0 || current != 0; {
		// 计算当前位1出现的个数并加到总结果
		if current == 0 {
			result += high * digit
		} else if current == 1 {
			result += high*digit + low + 1
		} else {
			result += (high + 1) * digit
		}

		// 向前移动当前位，并更新current、high、low、digit
		low += current * digit
		current = high % 10
		high /= 10
		digit *= 10
	}

	return result
}
