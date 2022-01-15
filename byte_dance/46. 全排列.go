package main

// 46. 全排列

// 给定一个不含重复数字的数组 nums ，返回其所有可能的全排列 。你可以 按任意顺序 返回答案。

// func main() {

// }

func permute(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	result := make([][]int, 0)
	dfs1(nums, 0, &result)
	return result
}

func dfs1(nums []int, index int, result *[][]int) {
	if index == len(nums) {
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
		dfs1(nums, index+1, result)
		nums[index], nums[i] = nums[i], nums[index]
	}
}
