package main

// 78. 子集

// func main() {

// }

// 排列组合，类似offer 38
// 回溯法
func subsets(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	res := make([]int, 0)      // 一个子集
	result := make([][]int, 0) // 所有子集结果
	for i := 0; i <= len(nums); i++ {
		dfs3(nums, 0, i, &res, &result)
	}
	return result
}

func dfs3(nums []int, start, num int, res *[]int, result *[][]int) {
	if len(*res) == num {
		tmp := make([]int, len(*res))
		copy(tmp, *res)
		*result = append(*result, tmp)
		return
	}

	for i := start; i < len(nums); i++ {
		*res = append(*res, nums[i])
		dfs3(nums, i+1, num, res, result)
		*res = (*res)[:len(*res)-1]
	}
}
