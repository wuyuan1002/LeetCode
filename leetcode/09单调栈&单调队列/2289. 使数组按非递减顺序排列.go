package main

// 2289. 使数组按非递减顺序排列

// 给你一个下标从 0 开始的整数数组 nums 。在一步操作中，移除所有满足 nums[i - 1] > nums[i] 的 nums[i] ，其中 0 < i < nums.length 。
//
// 重复执行步骤，直到 nums 变为 非递减 数组，返回所需执行的操作数。

// totalSteps .
// 每个数会将其右侧所有比自己小的数消除，直到碰到第一个不小于自己的数为止，也就是对于每个数都寻找其右侧首个不小于自己的元素
// 因为每一步操作，所有大于其右侧元素的元素都可以删除一次，也就是拥有最多右侧比其小的元素的数量就是最终要操作的步骤数
// 不能使用496那种正序遍历使用单调递增栈的方式，因为本题需要统计其左侧小的元素数量，需要有右侧更大值来触发计算，因此如[9,1,2,3]便不会被触发
// 因此本题应该倒过来，倒序遍历使用单调递增栈
func totalSteps(nums []int) int {
	// 用来记录每个元素能删掉的右侧元素数量
	type step struct {
		num   int
		count int
	}

	// 单调递增栈，用来存还没有出现下一个比其大的元素的下标
	stack := make([]step, 0)

	// 倒序遍历每个数字，若当前元素比栈顶元素小，则入栈，若当前元素比栈顶元素大，说明栈内比当前元素小的数字都是可以被当前元素删掉的
	for i := len(nums) - 1; i >= 0; i-- {
		// 当前元素能删掉的右侧数字数量最大值
		count := 0

		// 遍历栈内所有比当前元素小的元素，这些都是当前元素能删掉的元素
		for len(stack) > 0 && nums[i] > stack[len(stack)-1].num {
			// 获取栈顶元素并出栈
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 计算当前元素能删掉的右侧数字数量最大值
			count = max(count+1, top.count)
		}

		// 栈内已没有当前数字能删的元素了 -- 将当前元素与其能删掉的右侧数字数量入栈
		stack = append(stack, step{nums[i], count})
	}

	// 此时栈内存的是每个能删元素的数字及其能删掉的数字数量，这些里的最大值就是最终结果
	result := 0
	for _, s := range stack {
		result = max(result, s.count)
	}
	return result
}
