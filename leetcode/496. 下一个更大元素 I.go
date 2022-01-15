package main

// 496. 下一个更大元素 I

// 给你两个 没有重复元素 的数组nums1 和nums2，其中nums1是nums2的子集。
// 请你找出 nums1中每个元素在nums2中的下一个比其大的值。
// nums1中数字x的下一个更大元素是指x在nums2中对应位置的右边的第一个比x大的元素。
// 如果不存在，对应位置输出 -1 。

// func main() {

// }

// 一但要求下一个更大的元素，就是用单调栈解，leetcode相似的题目都是这个解法

// 类似739
// 同503，本题使用map来存一个列表中数字的下一个更大元素，再遍历下一个列表求下一个更大元素，
// 如果是求一个列表中数字的下一个更大元素则可以直接在res中保存，如503那样
// 单调栈 + map哈希表
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 首先倒序遍历nums2，维护一个单调递减栈，同时使用map记录下nums2中每个数字的下一个更大元素
	// 之后再遍历nums1，分别从map中获取对应数字的下一个更大元素
	stack := make([]int, 0)
	hash := make(map[int]int) // 保存元素与它的下一个更大元素
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		// 若栈中元素有小于当前元素的，把他们出栈，此时，栈中下一个元素即为当前num的下一个更大元素
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		// 记录当前num的下一个更大元素
		if len(stack) > 0 {
			hash[num] = stack[len(stack)-1]
		} else {
			hash[num] = -1
		}

		// 将当前元素入栈
		stack = append(stack, num)
	}

	// 遍历nums1，从map中挑出所需数字
	res := make([]int, len(nums1))
	for i, n := range nums1 {
		res[i] = hash[n]
	}

	return res
}
