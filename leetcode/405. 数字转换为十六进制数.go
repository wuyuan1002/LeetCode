package main

// 405. 数字转换为十六进制数

// 给定一个整数，编写一个算法将这个数转换为十六进制数。对于负整数，我们通常使用补码运算方法。
//
// 注意:
// 十六进制中所有字母(a-f)都必须是小写。
// 十六进制字符串中不能包含多余的前导零。如果要转化的数为0，那么以单个字符'0'来表示；对于其他情况，十六进制字符串中的第一个字符将不会是0字符。
// 给定的数确保在32位有符号整数范围内。
// 不能使用任何由库提供的将数字直接转换或格式化为十六进制的方法。

// func main() {
// 	fmt.Println(toHex(987))
// }

// 位运算
// 将num的二进制位进行分组，每4位为一组表示1～15的所有数字，将这些数字组成字符串返回
func toHex(num int32) string {
	if num == 0 {
		return "0"
	}

	// 反转字节数字
	reverse := func(res []byte) {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}

	unum := uint32(num) // 由于go没有无符号右移，所以先把num转成无符号int32
	res := make([]byte, 0)
	for ; unum != 0; unum >>= 4 {
		n := unum & 15
		if n < 10 {
			res = append(res, byte(n+'0'))
		} else {
			res = append(res, byte(n-10+'a'))
		}
	}

	// 反转结果并返回
	reverse(res)
	return string(res)
}
