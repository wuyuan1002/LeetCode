package main

// 17. 电话号码的字母组合

// func main() {

// }

var phone = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

// 回溯法
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	res := make([]byte, 0)
	result := make([]string, 0)
	dfs9(digits, 0, &res, &result)
	return result
}

func dfs9(digits string, index int, res *[]byte, result *[]string) {
	if index == len(digits) {
		*result = append(*result, string(*res))
		return
	}

	for _, num := range phone[digits[index]] {
		*res = append(*res, num)
		dfs9(digits, index+1, res, result)
		*res = (*res)[:len(*res)-1]
	}
}
