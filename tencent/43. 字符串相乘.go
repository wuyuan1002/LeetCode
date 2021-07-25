package main

import (
	"fmt"
	"strings"
)

// 43. 字符串相乘

// 给定两个以字符串形式表示的非负整数 num1 和 num2，
// 返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

func main() {
	fmt.Println(multiply("12", "3"))
}

// 类似 415 字符串相加
// 大数相乘
func multiply(num1 string, num2 string) string {
	if num1 == "" || num1 == "0" {
		return "0"
	} else if num2 == "" || num2 == "0" {
		return "0"
	}

	// 使用数组存储每次计算结果，避免大量的字符串相加运算
	res := make([]int, len(num1)+len(num2))
	for i := len(num2) - 1; i >= 0; i-- {
		// 取num2中每一位与num1中各数字相乘
		n1 := int(num2[i] - '0')
		for j := len(num1) - 1; j >= 0; j-- {
			n2 := int(num1[j] - '0')
			sum := res[i+j+1] + n1*n2

			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}

	str := strings.Builder{}
	for i := 0; i < len(res); i++ {
		if str.Len() == 0 && res[i] == 0 {
			// 跳过开头的0
			continue
		}
		str.WriteRune(rune(res[i] + '0'))
	}

	return str.String()
}
