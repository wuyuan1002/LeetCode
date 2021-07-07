package main

import (
	"fmt"
)

// 剑指 Offer 38. 字符串的排列

// 输入一个字符串，打印出该字符串中字符的所有排列。
// 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

func main() {
	fmt.Println(permutation("aab"))
	// findQueen()
}

// Hot100 46, 78
// 回溯法 + 剪枝
func permutation(s string) []string {
	result := make([]string, 0)
	if s == "" {
		return result
	}

	bytes := []byte(s)
	dfs(0, &bytes, &result)

	return result
}

// 第一步求所有可能出现在第一个位置的字符，即把第一个字符和后面所有的字符交换。
// 第二步固定第一个字符，求后面所有字符的排列。
func dfs(index int, bytes *[]byte, result *[]string) {
	if index == len(*bytes)-1 { // 若已经到最后一个字符，则将结果输出
		*result = append(*result, string(*bytes))
		return
	}

	visited := make(map[byte]bool) // 当前层是否出现过
	for i := index; i < len(*bytes); i++ {
		if visited[(*bytes)[i]] {
			continue // 若已经出现过，则直接跳过 -- 剪枝
		}

		visited[(*bytes)[i]] = true // 记录当前元素在当前层已经出现过 -- 已经被交换到第一个字符过

		(*bytes)[index], (*bytes)[i] = (*bytes)[i], (*bytes)[index] // 先将两元素交换，把可能出现在第一个的字符换过来
		dfs(index+1, bytes, result)                                 // 进入下一层
		(*bytes)[index], (*bytes)[i] = (*bytes)[i], (*bytes)[index] // 还原交换
	}
}

// -----------------

// 答案不对
// 八皇后问题 -- 求出皇后所有出现的位置的全排列，然后再挑选符合条件的
func findQueen() {
	columnIndex := make([]int, 8) // 下标为皇后所在的行号（下标一定不同，所以皇后一定不会处在同一行），下标对应的数字为皇后所在的列号
	for i := 0; i < len(columnIndex); i++ {
		columnIndex[i] = i // 使用0~7初始化列号，这样所有皇后一定不处在同一列
	}

	// 求出皇后所有出现的可能

	result := make([]*[]int, 0)
	dfsQueen(0, &columnIndex, &result)
	fmt.Println(len(result))
}

// index表示现在排列的是第几个皇后
func dfsQueen(index int, columnIndex *[]int, result *[]*[]int) {
	if index == len(*columnIndex)-1 { // 如果已经到最后一个了，看结果正不正确
		isTrue := true
		for i := 0; i < len(*columnIndex)-1; i++ {
			for j := i + 1; j < len(*columnIndex)-1; j++ {
				// 若任意两个皇后所在位置处于对角线上，则位置不正确
				if i-j == (*columnIndex)[i]-(*columnIndex)[j] || j-i == (*columnIndex)[i]-(*columnIndex)[j] {
					isTrue = false
					break
				}
			}
			if !isTrue {
				break
			}
		}

		// 若所有皇后位置正确，则天界到结果中
		if isTrue {
			*result = append(*result, columnIndex)
		}
		return
	}

	for i := 0; i < len(*columnIndex); i++ { // 求出皇后可以在的位置的所有排列
		(*columnIndex)[index], (*columnIndex)[i] = (*columnIndex)[i], (*columnIndex)[index] // 先将两元素交换，把可能出现在第一个的字符换过来
		dfsQueen(index+1, columnIndex, result)                                              // 进入下一层
		(*columnIndex)[index], (*columnIndex)[i] = (*columnIndex)[i], (*columnIndex)[index] // 还原交换
	}
}
