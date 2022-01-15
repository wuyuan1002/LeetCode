package main

// 66. 加一

// 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
// 你可以假设除了整数 0 之外，这个整数不会以零开头。

// func main() {
// 	fmt.Println(plusOne([]int{9, 9, 9, 9, 9}))
// }

// 将要加的数字用一个plus变量记录了下来，这样不只是加1可以，加任意数都可以计算了
func plusOne(digits []int) []int {
	if digits == nil || len(digits) == 0 {
		return []int{}
	}

	plus := 1 // 要加的数字
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += plus
		if digits[i] > 9 {
			if i > 0 {
				plus = digits[i] / 10
				digits[i] %= 10
			} else {
				digits[i] %= 10
				digits = append([]int{plus}, digits...)
			}
		} else {
			break
		}
	}
	return digits
}

// 只能进行加1操作
func plusOne1(digits []int) []int {
	if digits == nil || len(digits) == 0 {
		return []int{}
	}

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 { // 说明没有进位
			return digits
		}
		// 若当前位有进位，则继续遍历对前一个数字进行加1操作
	}
	digits = append([]int{1}, digits...)
	return digits
}
