package main

import "strconv"

// 738. 单调递增的数字

// 给定一个非负整数N，找出小于或等于N的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
//
// （当且仅当每个相邻位数上的数字x和y满足x <= y时，我们称这个整数是单调递增的。）

// monotoneIncreasingDigits .
// 从后向前遍历，一旦出现num[i-1] > num[i]，就让num[i-1]--，同时num[i] = 9
func monotoneIncreasingDigits(n int) int {
	num := []byte(strconv.Itoa(n))
	flag := len(num) // flag之后的数字全部变成9，标记从哪里之后全部赋值为9 -- 如 例子：239586 -> 238999

	// 从后向前遍历每一位，遇到前一位比后一位大的，向前一位借1，后一位置为9
	for i := len(num) - 1; i > 0; i-- {
		if num[i-1] > num[i] {
			// 向前一位借1
			num[i-1]--
			// 标记当前位置之后的全部置为9
			flag = i
		}
	}

	// 将flag后的所有位都置为9
	for i := flag; i < len(num); i++ {
		num[i] = '9'
	}

	// 字符串转成数字
	n, _ = strconv.Atoi(string(num))

	return n
}
