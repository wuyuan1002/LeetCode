package main

import "fmt"

// 47. 全排列 II

// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

// 类似 46
func permuteUnique(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	result := make([][]int, 0)
	dfs5(nums, 0, &result)
	return result
}

func dfs5(nums []int, index int, result *[][]int) {
	if index == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
		return
	}

	visited := make(map[int]bool)
	for i := index; i < len(nums); i++ {
		if visited[nums[i]] {
			continue
		}

		visited[nums[i]] = true
		nums[index], nums[i] = nums[i], nums[index]
		dfs5(nums, index+1, result)
		nums[index], nums[i] = nums[i], nums[index]
	}
}
