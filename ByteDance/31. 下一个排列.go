package main

// 31. 下一个排列

// 实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
// 必须 原地 修改，只允许使用额外常数空间。

func main() {

}

func nextPermutation(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	// 从后往前找第一个小于后一个元素的下标
	k := len(nums) - 2
	for ; k >= 0; k-- {
		if nums[k] < nums[k+1] {
			break
		}
	}

	// 若不存在，则反转整个数组 -- 此时说明数组已经是倒序排列了，最大
	if k < 0 {
		reverseArr(nums)
		return
	}

	// 从后往前找第一个大于nums[k]的元素的下标
	l := len(nums) - 1
	for ; l >= 0; l-- {
		if nums[l] > nums[k] {
			break
		}
	}

	// 交换nums[k]和nums[l]
	nums[k], nums[l] = nums[l], nums[k]

	// 翻转k之后的元素
	reverseArr(nums[k+1:])
}

// 翻转数组
func reverseArr(nums []int) {
	if nums == nil || len(nums) <= 1 {
		return
	}

	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
