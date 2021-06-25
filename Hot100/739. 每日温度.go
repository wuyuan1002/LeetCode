package main

import (
	"fmt"
)

// 739. 每日温度

// 请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。
// 如果气温在这之后都不会升高，请在该位置用 0 来代替。
// 例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，
// 你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
//
// 提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

// 单调栈 -- 单调递减栈
// 见Hot100 42,84
func dailyTemperatures(temperatures []int) []int {
	if temperatures == nil || len(temperatures) == 0 {
		return nil
	}

	result := make([]int, len(temperatures))

	stack := make([]int, 0)
	for i := 0; i < len(temperatures); i++ {

		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result[top] = i - top
		}

		stack = append(stack, i)
	}

	return result
}
