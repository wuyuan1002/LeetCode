package main

// 69. x的平方根

// 实现int sqrt(int x)函数。
// 计算并返回x的平方根，其中x是非负整数。
//
// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

func main() {

}

// 二分查找
func mySqrt(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}

	res := -1
	l, r := 0, x
	for l <= r {
		mid := (l + r) / 2
		if mid*mid <= x {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}
