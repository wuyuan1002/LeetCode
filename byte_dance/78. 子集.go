package main

// 78. 子集

// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

// func main() {
// 	fmt.Println(subsets([]int{1, 2, 2}))
// }

// 排列组合，类似offer 38
// 回溯法
func subsets(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	res := make([]int, 0)
	result := make([][]int, 0)
	for i := 0; i <= len(nums); i++ {
		dfs9(nums, 0, i, &res, &result)
	}

	return result
}

// index: 从哪个数字开始
// num: 这次要组合几个数字
func dfs9(nums []int, index int, num int, res *[]int, result *[][]int) {
	if len(*res) == num {
		tmp := make([]int, num)
		copy(tmp, *res)
		*result = append(*result, tmp)
		return
	}

	for i := index; i < len(nums); i++ {
		*res = append(*res, nums[i])
		dfs9(nums, i+1, num, res, result)
		*res = (*res)[:len(*res)-1]
	}
}
