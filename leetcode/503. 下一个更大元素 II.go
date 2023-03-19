package main

// 503. 下一个更大元素 II

// 给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），
// 输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，
// 这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。
// 如果不存在，则输出 -1。

// func main() {
// 	fmt.Println(nextGreaterElements([]int{2, 5, 1, 4}))
// }

// 单调栈
// 类似739
// 同496，只是本题还需要考虑循环数组的问题，如何实现循环数组呢?
// -> 直接将原数组复制一份拼接到数组后面,只是在实现时不直接这样做而是通过将下标取模来实现
//
// 保存下一个更大元素时，496是先使用map保存之后再求，也可以直接保存到res中
// 由于要知道当前元素的下标，所以在stack中保存的是数字的下标
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	stack := make([]int, 0) // 单调递减栈
	res := make([]int, n)   // 结果集
	for i := range res {
		res[i] = -1 // 初始化res各个元素的下标为-1，先假设每个元素都没有下一个更大元素
	}

	// 遍历nums
	for i := 0; i < n*2-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			// 若栈顶元素小于当前元素，则栈顶元素的下一个更大元素就是当前元素
			res[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}

	return res
}
