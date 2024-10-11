package main

// 67. 二进制求和

// 给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。

// addBinary .
// leetcode 415. 字符串相加
// leetcode 43. 字符串相乘
//
// 此题与415完全一致，只是415是十进制，此题是二进制（在求当前位得数和进位时使用的是2而不是10）
func addBinary(a string, b string) string {
	result := make([]byte, 0)  // 总结果
	carry := 0                 // 进位
	i, j := len(a)-1, len(b)-1 // 两个指针分别指向两数末尾

	// 不断计算每一位的得数和进位，并将每一位的得数添加到总结果尾部
	for i >= 0 || j >= 0 || carry > 0 {
		// 获取i、j位置对应的的数字
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(a[i] - '0')
		}
		if j >= 0 {
			n2 = int(b[j] - '0')
		}

		// 计算当前位得数和进位
		sum := n1 + n2 + carry
		result = append(result, byte(sum%2)+'0')
		carry = sum / 2

		// 同时向前移动两指针
		i--
		j--
	}

	// 反转总结果切片
	reverseSlice(result)

	return string(result)
}

// reverseSlice 反转切片
func reverseSlice(nums []byte) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
