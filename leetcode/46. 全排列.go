package main

// 46. 全排列

// func main() {

// }

// offer 38
// 回溯法+剪枝
func permute(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	result := make([][]int, 0)
	dfs18(nums, 0, &result)
	return result
}

func dfs18(nums []int, index int, result *[][]int) {
	if index == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
		return
	}

	visited := make(map[int]bool) // 剪枝 -- 去重
	for i := index; i < len(nums); i++ {
		if visited[nums[i]] {
			continue
		}
		visited[nums[i]] = true

		nums[index], nums[i] = nums[i], nums[index]
		dfs18(nums, index+1, result)
		nums[index], nums[i] = nums[i], nums[index]
	}
}
