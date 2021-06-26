package main

// 46.全排列 -- 排列

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

func main() {
	permute([]int{1, 2, 3})
}

// 排列组合，类似offer 38, Hot100 78
// 回溯法+剪枝
func permute(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	result := make([][]int, 0)
	dfs3(nums, 0, &result)
	return result
}

func dfs3(nums []int, index int, result *[][]int) {
	if index == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)

		return
	}

	visited := make(map[int]bool) // 当前层是否出现过

	for i := index; i < len(nums); i++ {
		if visited[nums[i]] {
			continue
		}

		visited[nums[i]] = true

		nums[index], nums[i] = nums[i], nums[index]
		dfs3(nums, index+1, result)
		nums[index], nums[i] = nums[i], nums[index]
	}
}
