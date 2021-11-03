package main

// 67. 二进制求和

// 给你两个二进制字符串，返回它们的和（用二进制表示）。
//
// 输入为 非空 字符串且只包含数字 1 和 0。

func main() {

}

// 类似 415大数相加
func addBinary(a string, b string) string {
	res := make([]byte, 0)

	carry := 0     // 进位
	n1, n2 := 0, 0 // 两数在i,j位置的数字
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		if i >= 0 {
			n1 = int(a[i] - '0')
		} else {
			n1 = 0
		}

		if j >= 0 {
			n2 = int(b[j] - '0')
		} else {
			n2 = 0
		}

		sum := n1 + n2 + carry             // 当前位之和
		carry = sum / 2                    // 进位
		res = append(res, byte(sum%2)+'0') // 当前位结果
	}

	// 翻转结果数组
	reverseString(res)

	return string(res)
}
