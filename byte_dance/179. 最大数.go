package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 179. 最大数

// 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

func main() {
	nums := []int{0, 0}
	fmt.Println(largestNumber(nums))
}

// 类似 Offer 45 把数组排成最小的数
func largestNumber(nums []int) string {
	if nums == nil || len(nums) == 0 {
		return ""
	}

	strArr := make([]string, len(nums))
	for i, num := range nums {
		strArr[i] = strconv.Itoa(num)
	}

	for i := 1; i < len(strArr); i++ {
		j := i - 1
		tmp := strArr[i]
		for ; j >= 0; j-- {
			if tmp+strArr[j] > strArr[j]+tmp {
				strArr[j+1] = strArr[j]
			} else {
				break
			}
		}
		strArr[j+1] = tmp
	}

	res := strings.Join(strArr, "")
	if n, _ := strconv.Atoi(res); n == 0 {
		// 排除输入为[0,0,0]却返回"000"的情况
		return "0"
	}
	return res
}
