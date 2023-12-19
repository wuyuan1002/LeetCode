package main

// 415. 字符串相加

// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
//
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。

// addStrings .
// 两指针分别从两字符串尾部进行相加操作，同时使用一个变量记录进位，将进位相加到下一位的计算中，
// 使用一个切片保存每一位的得数，将每一位的计算结果添加到切片末尾，最终计算完后反转整个切片即为最终结果
func addStrings(num1 string, num2 string) string {
	result := make([]byte, 0)        // 总结果
	carry := 0                       // 进位
	i, j := len(num1)-1, len(num2)-1 // 两个指针分别指向两数末尾

	for i >= 0 || j >= 0 || carry > 0 {
		// 获取i、j位置对应的的数字
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		}

		// 计算当前位得数和进位
		sum := n1 + n2 + carry
		result = append(result, byte((sum%10)+'0'))
		carry = sum / 10

		// 同时向前移动两指针
		i--
		j--
	}

	// 反转总结果切片
	reverse415(result)

	return string(result)
}

// reverse415 .
func reverse415(nums []byte) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
