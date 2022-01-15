package main

// 739. 每日温度

// 请根据每日 气温 列表 temperatures ，请计算在每一天需要等几天才会有更高的温度。
// 如果气温在这之后都不会升高，请在该位置用 0 来代替。

// func main() {

// }

// 单调栈 -- 单调递减栈
func dailyTemperatures(temperatures []int) []int {
	if temperatures == nil || len(temperatures) == 0 {
		return nil
	}

	res := make([]int, len(temperatures))

	stack := make([]int, 0)
	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res[top] = i - top
		}
		stack = append(stack, i)
	}

	return res
}

// 单调栈 -- 单调递减栈
func dailyTemperatures1(temperatures []int) []int {
	if temperatures == nil || len(temperatures) == 0 {
		return nil
	}

	res := make([]int, len(temperatures))

	stack := make([]int, 0)
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res[top] = i - top
		}
		stack = append(stack, i)
	}

	return res
}
