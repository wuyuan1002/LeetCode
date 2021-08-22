package main

import "strconv"

// 93. 复原 IP 地址

func main() {

}

// 回溯法
func restoreIpAddresses(s string) []string {
	if s == "" {
		return nil
	}
	if _, err := strconv.Atoi(s); err != nil {
		return nil
	}

	result := make([]string, 0)
	res := make([]string, 0)
	dfs16(s, 0, &res, &result)
	return result
}

func dfs16(s string, start int, res *[]string, result *[]string) {
	if len(*res) == 4 {
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

	for i := start; i < len(s); i++ {
		str := s[start : i+1]
		if str != "0" && str[0] == '0' {
			return
		} else if n, _ := strconv.Atoi(str); n > 255 {
			return
		}

		*res = append(*res, str)
		dfs16(s, i+1, res, result)
		*res = (*res)[:len(*res)-1]
	}
}
