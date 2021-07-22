package main

import (
	"fmt"
)

// 415. 字符串相加

// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

func main() {
	fmt.Println(addStrings("347", "20"))
}

// 类似 43. 字符串相乘
// 大数相加
func addStrings(num1 string, num2 string) string {
	if num1 == "" {
		return num2
	} else if num2 == "" {
		return num1
	}

	res := make([]byte, 0)           // 两数之和
	carry := 0                       // 进位
	i, j := len(num1)-1, len(num2)-1 // 两个指针分贝指向两数末尾

	n1, n2 := 0, 0 // 两数在i,j位置的数字
	for i >= 0 || j >= 0 || carry > 0 {
		if i >= 0 {
			n1 = int(num1[i] - '0')
		} else {
			n1 = 0
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		} else {
			n2 = 0
		}

		sum := n1 + n2 + carry                // 当前位的和
		res = append(res, byte((sum)%10)+'0') // 当前位
		carry = (sum) / 10                    // 进位

		i--
		j--
	}

	// 翻转结果数组
	revers(res)

	return string(res)
}

func revers(nums []byte) {
	if nums == nil || len(nums) == 0 {
		return
	}

	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
