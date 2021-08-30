package main

import "strconv"

// 738. 单调递增的数字

// 给定一个非负整数N，找出小于或等于N的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
//
// （当且仅当每个相邻位数上的数字x和y满足x <= y时，我们称这个整数是单调递增的。）

func main() {

}

// 从后向前遍历，一旦出现num[i-1] > num[i]，就让num[i-1]--，同时num[i] = 9
func monotoneIncreasingDigits(n int) int {
	num := []byte(strconv.Itoa(n))
	flag := len(num)
	for i := len(num) - 1; i > 0; i-- {
		if num[i-1] > num[i] {
			num[i-1]--
			num[i] = '9'

			// flag之后的数字全部变成9 -- 如 例子：1000
			flag = i
		}
	}

	for i := flag; i < len(num); i++ {
		num[i] = '9'
	}

	n, _ = strconv.Atoi(string(num))
	return n
}
