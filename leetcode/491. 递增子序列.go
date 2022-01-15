package main

// 491. 递增子序列

// 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，
// 递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
//
// 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。

// func main() {

// }

// func findSubsequences(nums []int) [][]int {
// 	if nums == nil || len(nums) < 2 {
// 		return nil
// 	}
//
// 	res := make([]int, 0)
// 	result := make([][]int, 0)
// 	dfs17(nums, 0, &res, &result)
// 	return result
// }
//
// func dfs17(nums []int, start int, res *[]int, result *[][]int) {
// 	if len(*res) >= 2 {
// 		tmp := make([]int, len(*res))
// 		copy(tmp, *res)
// 		*result = append(*result, tmp)
// 	}
//
// 	visited := make(map[int]bool)
// 	for i := start; i < len(nums); i++ {
// 		if len(*res) > 0 && nums[i] < (*res)[len(*res)-1] {
// 			continue
// 		}
//
// 		if visited[nums[i]] {
// 			continue
// 		}
//
// 		visited[nums[i]] = true
//
// 		*res = append(*res, nums[i])
// 		dfs17(nums, i+1, res, result)
// 		*res = (*res)[:len(*res)-1]
// 	}
// }

// 同 90子集，相当于只求递增的子集
func findSubsequences(nums []int) [][]int {
	if nums == nil || len(nums) < 2 {
		return nil
	}

	res := make([]int, 0)
	result := make([][]int, 0)
	for i := 2; i <= len(nums); i++ {
		dfs17(nums, 0, i, &res, &result)
	}
	return result
}

func dfs17(nums []int, start, n int, res *[]int, result *[][]int) {
	if len(*res) == n {
		tmp := make([]int, len(*res))
		copy(tmp, *res)
		*result = append(*result, tmp)
		return
	}

	visited := make(map[int]bool) // 由于数组未排序，所以使用map去重，而不是像90那样hi姐比较前一个数字
	for i := start; i < len(nums); i++ {
		if len(*res) > 0 && nums[i] < (*res)[len(*res)-1] {
			continue
		}

		if visited[nums[i]] {
			continue
		}
		visited[nums[i]] = true

		*res = append(*res, nums[i])
		dfs17(nums, i+1, n, res, result)
		*res = (*res)[:len(*res)-1]
	}
}
