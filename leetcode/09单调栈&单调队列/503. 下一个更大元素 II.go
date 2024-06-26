package main

// 503. 下一个更大元素 II

// 给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），
// 输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，
// 这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。
// 如果不存在，则输出 -1。

// nextGreaterElements .
// 解题方法和496一样，本题是求数组中每一个元素的下一个更大元素，
// 我们可以将数组首位相连后组成一个新数组，然后整体遍历一遍新数组，
// 使用map存起来每个元素的下一个更大元素，然后再构造总结果
//
// 也可以直接使用下标的方式更新结果集，
// 因为要使用栈内元素直接定位到在结果里的位置，所以栈内存的是元素下标而不是元素本身的值
//
// 本题需要考虑循环数组的问题，如何实现循环数组呢？
// -> 直接将原数组复制一份拼接到数组后面,只是在实现时不直接这样做而是通过将下标取模来实现
func nextGreaterElements(nums []int) []int {
	result := make([]int, len(nums)) // 结果集
	for i := range result {
		result[i] = -1 // 初始化res各个元素的下标为-1，先假设每个元素都没有下一个更大元素
	}

	stack := make([]int, 0) // 单调递减栈，用来存还没有出现下一个比其大的元素的下标

	// 因为本题是一个循环数组，所以将nums首尾相连，相当于遍历两遍nums数组
	for i := 0; i < len(nums)*2-1; i++ {
		// 若当前元素比栈顶元素大，开始不断将栈顶元素出栈，并记录栈顶出栈元素的下一个更大元素为当前元素到结果集
		for len(stack) > 0 && nums[i%len(nums)] > nums[stack[len(stack)-1]] {
			// 将栈顶元素出栈
			numIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 记录栈顶元素的下一个更大元素为当前元素
			// 因为这里要使用栈内元素直接定位到在结果里的位置，所以栈内存的是元素下标而不是元素本身的值
			result[numIndex] = nums[i%len(nums)]
		}

		// 栈内剩下的元素都是大于等于当前元素的元素了
		// 将当前元素下标入栈，等待自己的下一个更大元素出现后记录到结果集中
		stack = append(stack, i%len(nums))
	}

	return result
}
