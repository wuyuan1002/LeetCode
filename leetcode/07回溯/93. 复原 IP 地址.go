package main

import "strconv"

// 93. 复原 IP 地址

// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，
// 但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，
// 这些地址可以通过在 s 中插入'.' 来形成。
// 你 不能重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

// restoreIpAddresses .
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)    // 存放一次回溯路径分割的数字
	result := make([]string, 0) // 总结果集
	dfsRestoreIpAddresses(s, 0, &res, &result)
	return result
}

// dfsRestoreIpAddresses .
// s: 选择列表
// start: 指定当前层的选择范围 -- s[start: ]
// res: 存放一次回溯路径分割的数字
// result: 总结果 -- 分割后符合条件的IP列表
func dfsRestoreIpAddresses(s string, start int, res *[]string, result *[]string) {
	// 若已选择足够数量的符合IP地址的数字 -- 没必要进入下一层继续选择第5个数字了, 检查一下当前回溯路径的数字是否符合IP地址
	if len(*res) == 4 {
		// 若整个字符串都已经遍历到了, 并且恰好分成了4个数字, 说明切分到了一个符合条件的IP地址 -- 加入结果集
		// 否则说明没遍历完整个字符串时就分割够了4个数字, 不符合条件 -- 直接返回
		if start == len(s) {
			str := ""
			for i, n := range *res {
				str += n
				if i < len(*res)-1 {
					str += "."
				}
			}
			*result = append(*result, str)
		}
		return
	}

	// 不断从本层分割数字放入回溯路径然后进入下一层继续选择下一个数字
	for i := start; i < len(s); i++ {
		// 本次分割的字符串
		str := s[start : i+1]

		// 剪枝 -- 本次分割到的字符串不满足IP地址的数字要求 -- 已经不满足条件了, 直接返回即可
		if str != "0" && str[0] == '0' {
			return
		} else if n, _ := strconv.Atoi(str); n > 255 {
			return
		}

		// 将当前分割到的字符串加入回溯路径
		*res = append(*res, str)
		// 进入下一层选择下一个数字
		dfsRestoreIpAddresses(s, i+1, res, result)
		// 将当前数字移出回溯路径, 从当前层分割另一个数字
		*res = (*res)[:len(*res)-1]
	}
}
