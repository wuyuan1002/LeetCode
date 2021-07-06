package main

import (
	"fmt"
)

// 39. 组合总和

// 给定一个无重复元素的数组candidates和一个目标数target，
// 找出candidates中所有可以使数字和为target的组合。
//
// candidates中的数字可以无限制重复被选取。
//
// 说明：
//
// 所有数字（包括target）都是正整数。
// 解集不能包含重复的组合。

func main() {
	fmt.Println(combinationSum([]int{2, 6, 3, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	if candidates == nil || len(candidates) == 0 {
		return nil
	}

	quickSort1(candidates, 0, len(candidates)-1)

	result := make([][]int, 0)
	res := make([]int, 0)
	dfs5(candidates, 0, target, 0, &res, &result)

	return result
}

func dfs5(nums []int, currentSum, targetSum, start int, res *[]int, result *[][]int) {
	if currentSum == targetSum {
		tmp := make([]int, len(*res))
		copy(tmp, *res)
		*result = append(*result, tmp)
		return
	}

	for i := start; i < len(nums); i++ {
		// 剪枝
		if currentSum+nums[i] > targetSum {
			return
		}

		*res = append(*res, nums[i])
		dfs5(nums, currentSum+nums[i], targetSum, i, res, result)
		*res = (*res)[:len(*res)-1]
	}
}

func quickSort1(arr []int, left, right int) {
	if arr == nil || len(arr) == 0 || left >= right {
		return
	}

	l, r := left, right
	tmp := arr[left]

	for l < r {
		for l < r && arr[r] >= tmp {
			r--
		}
		for l < r && arr[l] <= tmp {
			l++
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	arr[l], arr[left] = tmp, arr[l]

	quickSort1(arr, left, l-1)
	quickSort1(arr, r+1, right)
}
