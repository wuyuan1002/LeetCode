package main

import (
	"fmt"
)

// 39. 组合总和

// 给定一个无重复元素的数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
// candidates中的数字可以无限制重复被选取。
//
// 说明：
// 所有数字（包括target）都是正整数。
// 解集不能包含重复的组合。

func main() {

	// aa := []int{3, 5, 6, 1, 2, 8, 3, 4}
	// sort1(aa, 0, len(aa)-1)
	// fmt.Println(aa)

	aa := combinationSum([]int{3, 5, 6, 1, 2, 8, 3, 4}, 7)
	fmt.Println(aa)

}

// 回溯法+剪枝
func combinationSum(candidates []int, target int) [][]int {
	if candidates == nil || len(candidates) == 0 {
		return nil
	}

	// 先排序数组，方便剪枝
	sort1(candidates, 0, len(candidates)-1)

	// 回溯求解
	result := make([][]int, 0)
	res := make([]int, 0)
	dfs2(candidates, 0, target, 0, &res, &result)

	return result
}

// candidates: 总数组
// currentSum: 上一层的和
// target: 目标值
// start: 从哪个位置开始当前层的数字遍历
// res: 路径
// result: 总结果
func dfs2(candidates []int, currentSum, target int, start int, res *[]int, result *[][]int) {

	// 路径结束条件 -- 若总和已经等于目标值，将路径粗存入最终结果中并返回
	if currentSum == target {
		tmp := make([]int, len(*res))
		copy(tmp, *res)
		*result = append(*result, tmp)

		return
	}

	num := 0 // 当前数字
	for i := start; i < len(candidates); i++ {

		num = candidates[i]

		// 剪枝 -- 由于数组已排序，若加上当前数字都已经比目标值大了，那么加上后面的数字更大
		if currentSum+num > target {
			break
		}

		// 将当前数字加到路径中然后
		*res = append(*res, num)
		dfs2(candidates, currentSum+num, target, i, res, result)
		*res = (*res)[:len(*res)-1]
	}
}

func sort1(arr []int, left, right int) {
	if left >= right || arr == nil || len(arr) == 0 {
		return
	}

	l, r := left, right
	tmp := arr[left]

	for l < r {
		for arr[r] >= tmp && l < r {
			r--
		}
		for arr[l] <= tmp && l < r {
			l++
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	arr[l], arr[left] = tmp, arr[l]

	sort1(arr, left, l-1)
	sort1(arr, l+1, right)
}
