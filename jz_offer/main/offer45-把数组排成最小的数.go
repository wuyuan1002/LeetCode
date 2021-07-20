package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 剑指 Offer 45. 把数组排成最小的数

// 输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

func main() {
	nums := []int{3, 30, 34, 5, 9}
	fmt.Println(minNumber(nums))
}

func minNumber(nums []int) string {
	if nums == nil || len(nums) == 0 {
		return ""
	}
	var strNums []string
	for _, num := range nums {
		strNums = append(strNums, strconv.Itoa(num))
	}
	// sort.Slice(strNums, func(i, j int) bool {
	// 	if strNums[i]+strNums[j] < strNums[j]+strNums[i] {
	// 		return true
	// 	}
	// 	return false
	// })

	j := 0
	temp := ""
	for i := 1; i < len(strNums); i++ { // 类似插入排序
		temp = strNums[i]
		for j = i - 1; j >= 0; j-- {
			if temp+strNums[j] < strNums[j]+temp {
				strNums[j+1] = strNums[j]
				continue
			}
			break
		}
		strNums[j+1] = temp
		// for j, l = i-1, i; j >= 0; {
		// 	if strNums[l]+strNums[j] < strNums[j]+strNums[l] {
		// 		strNums[l], strNums[j] = strNums[j], strNums[l]
		// 		j--
		// 		l--
		// 		continue
		// 	}
		// 	break
		// }
	}
	return strings.Join(strNums, "")
}

// 第二次做 -- 插入排序，把数字数组转换成string数组，之后遍历数组。比较字符串拼接后的字典顺序，同时可以解决大数问题
func minNumber1(nums []int) string {
	if nums == nil || len(nums) == 0 {
		return ""
	}

	// 将数字数组转换成字符串数组
	strArr := make([]string, len(nums))
	for i, num := range nums {
		strArr[i] = strconv.FormatInt(int64(num), 10)
	}

	// 类似于插入排序，将数字插入到合适的位置
	for i := 1; i < len(strArr); i++ {
		j := i - 1
		temp := strArr[i]
		for ; j >= 0; j-- {
			if temp+strArr[j] < strArr[j]+temp {
				strArr[j+1] = strArr[j]
				continue
			}
			break
		}
		strArr[j+1] = temp
	}

	res := strings.Join(strArr, "")
	if n, _ := strconv.Atoi(res); n == 0 {
		// 排除输入为[0,0,0]却返回"000"的情况
		return "0"
	}
	return res
}
