package main

import (
	"fmt"
	"strconv"
)

// 93. 复原 IP 地址

// 给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。
// 你可以按任何顺序返回答案。
//
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，
// 但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。

func main() {
	// fmt.Println(restoreIpAddresses("101023"))
	// fmt.Println(restoreIpAddresses("010010"))
	fmt.Println(restoreIpAddresses("25525511135"))
}

// 回溯法
func restoreIpAddresses(s string) []string {
	if s == "" {
		return nil
	} else if _, err := strconv.Atoi(s); err != nil {
		// IP必须由数字构成，不能有别的符号
		return nil
	}

	res := make([]string, 0)
	result := make([]string, 0)
	dfs7(s, &res, &result)

	return result
}

// res: 存IP的每一个数字
// result: 存每一个IP
func dfs7(str string, res *[]string, result *[]string) {
	if str == "" || len(*res) == 4 {
		if str == "" && len(*res) == 4 { // IP地址必须是4位数字
			ip := ""
			for i, num := range *res { // 遍历每一个数字，拼接IP
				if num != "0" && num[0] == '0' {
					// 若当前数字不是0，但以0开头，过滤掉不合法
					return
				} else if n, _ := strconv.Atoi(num); n > 255 {
					// IP的每一位数字不能大于255
					return
				}

				ip += num
				if i != len(*res)-1 {
					ip += "."
				}
			}
			*result = append(*result, ip)
		}
		return
	}

	for i := 1; i <= len(str); i++ {
		*res = append(*res, str[:i])
		dfs7(str[i:], res, result)
		*res = (*res)[:len(*res)-1]
	}
}
