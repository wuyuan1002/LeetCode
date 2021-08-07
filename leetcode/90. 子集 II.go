package main

import "fmt"

// 90. 子集 II

// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	sort(nums, 0, len(nums)-1)

	res := make([]int, 0)      // 一个子集
	result := make([][]int, 0) // 所有子集结果
	for i := 0; i <= len(nums); i++ {
		dfs4(nums, 0, i, &res, &result)
	}
	return result
}

func dfs4(nums []int, start, num int, res *[]int, result *[][]int) {
	if len(*res) == num {
		tmp := make([]int, len(*res))
		copy(tmp, *res)
		*result = append(*result, tmp)
		return
	}

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			// 剪枝，防止出现重复子集
			continue
		}
		*res = append(*res, nums[i])
		dfs4(nums, i+1, num, res, result)
		*res = (*res)[:len(*res)-1]
	}
}

func sort(nums []int, left, right int) {
	if nums == nil || len(nums) == 0 || left >= right {
		return
	}

	l, r := left, right
	tmp := nums[left]
	for l < r {
		for l < r && nums[r] >= tmp {
			r--
		}
		for l < r && nums[l] <= tmp {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[l], nums[left] = tmp, nums[l]
	sort(nums, left, l-1)
	sort(nums, r+1, right)
}
