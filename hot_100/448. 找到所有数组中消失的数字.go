package main

// 448. 找到所有数组中消失的数字

// 给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。
// 请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。
//
// 进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。

// func main() {

// }

// 遍历到一个数字，都把数字-1对应到的下标+n，之后再次遍历一次数组，数字大于n的位置的下标+1就是没出现过的数字
// [4,3,2,7,8,2,3,1]，4改变了下标为3的数字变成了nums[3] + 8 = 15;
// 那么遍历到15时，(15-1)%8=6 和 (7-1)%8=6是一致的，不会影响结果，
// 所以才要加上n，当然只要是n的倍数都可以。模n是必不可少的，因为被改变了的数一定大于n。
func findDisappearedNumbers(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	result := make([]int, 0)

	n := len(nums)
	for _, num := range nums {
		v := (num - 1) % n
		nums[v] += n
	}

	for i := range nums {
		if nums[i] <= n {
			result = append(result, i+1)
		}
	}

	return result
}
