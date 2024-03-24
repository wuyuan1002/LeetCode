package main

import (
	"strconv"
	"strings"
)

// 179. 最大数

// 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

// largestNumber .
// 把数字数组转换成string数组，之后将数组进行排序，然后将数组中数字全部进行拼接即可，同时可以解决大数问题
// 只是比较数字大小时不是直接比较两数字原本的大小，而是比较两数字字符串拼接后的字典顺序，
// 如比较"7"和"38"，由于"738" > "387"，所以"7"大于"38"，排序时要排在前面，
func largestNumber(nums []int) string {

	// 将数字数组转换成字符串数组
	strArr := make([]string, len(nums))
	for i, num := range nums {
		strArr[i] = strconv.Itoa(num)
	}

	// 按照两数字字符串拼接后的数字字典顺序进行排序
	// 此处使用快排对数组进行排序
	quickSort179(strArr, 0, len(strArr)-1)

	// 若数组第一个数字为0，说明整体为0 -- 排除输入为[0,0,0]却返回"000"的情况
	if len(strArr) > 0 && strArr[0] == "0" {
		return "0"
	}

	// 将数组中字符串拼接并返回
	return strings.Join(strArr, "")
}

// quickSort179 将数组按照前后数字字典序从大到小排序
func quickSort179(nums []string, left, right int) {
	if left >= right {
		return
	}

	l, r := left, right
	temp := nums[l]

	for l < r {
		for l < r && nums[r]+temp <= temp+nums[r] {
			r--
		}
		for l < r && nums[l]+temp >= temp+nums[l] {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

	nums[left], nums[l] = nums[l], nums[left]

	quickSort179(nums, left, l-1)
	quickSort179(nums, r+1, right)
}
