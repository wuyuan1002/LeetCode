package main

// 167. 两数之和 II - 输入有序数组

// 给定一个已按照 升序排列 的整数数组numbers ，请你从数组中找出两个数满足相加之和等于目标数target 。
//
// 函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 1 开始计数 ，
// 所以答案数组应当满足 1 <= answer[0] < answer[1] <= numbers.length 。
//
// 你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。

func main() {

}

// 双指针
func twoSum1(numbers []int, target int) []int {
	if numbers == nil || len(numbers) == 0 {
		return nil
	}

	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}

	return nil
}
