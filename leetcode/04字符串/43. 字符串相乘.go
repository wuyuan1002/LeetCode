package main

import "strings"

// 43. 字符串相乘

// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
// 注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。

// multiply .
// 同 leetcode 415. 字符串相加、67. 二进制求和
//
// 遍历 num2 的每一位与 num1 进行相乘，将每一步的结果进行累加
func multiply(num1 string, num2 string) string {
	if num1 == "" || num1 == "0" {
		return "0"
	}
	if num2 == "" || num2 == "0" {
		return "0"
	}

	// 使用数组存储每次计算结果，避免大量的字符串相加运算 -- 数组的每一位代表结果的每一位
	res := make([]int, len(num1)+len(num2))

	// 取num2中每一位与num1中各数字相乘
	for i := len(num2) - 1; i >= 0; i-- {
		// 获取当前位num2的数字
		n1 := int(num2[i] - '0')
		// 取num1中每一位数字与当前位num2的数字相乘
		for j := len(num1) - 1; j >= 0; j-- {
			// 获取当前位num1的数字
			n2 := int(num1[j] - '0')
			// 计算当前位的得数 -- res[i+j+1]存的是前一位计算的进位
			sum := res[i+j+1] + n1*n2

			// 保存当前位和进位的得数
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}

	// 把最终的int数组转换成字符串
	str := strings.Builder{}
	for _, num := range res {
		if str.Len() == 0 && num == 0 {
			// 跳过开头的0
			continue
		}
		str.WriteByte(byte(num + '0'))
	}

	return str.String()
}
