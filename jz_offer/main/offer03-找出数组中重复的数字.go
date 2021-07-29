package main

import "fmt"

// 剑指 Offer 03. 数组中重复的数字

// 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
// 数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
// 请找出数组中任意一个重复的数字。

func main() {
	nums := []int{2, 2, 2, 2, 2, 5, 5}
	fmt.Printf("%d", findRepeatNumber2(nums))
}

// 因为每个数字都在0~n-1范围内，所以如果所有数字都不重复的话，每个数字对应一个下标，但是若有数字重复的话
// 就会有一个数字对应多个下标。因此，可以遍历数组，把遇到的每一个数字都放到它对应的下标处，如果在放某个数字时发现
// 下标处的数字已经和下标相等了，那么这个数字就是重复数字
func findRepeatNumber(nums []int) int {
	i := 0  // 下标
	n := -1 // 寄存当前下标i的数字
	for i < len(nums) {
		n = nums[i]
		if n >= len(nums) {
			panic("数组中数字不能大于数组下标")
		}
		if n != i {
			if n != nums[n] {
				nums[n], nums[i] = nums[i], nums[n]
				continue
			}
			break
		}
		i++
		n = -1
	}
	return n
}

// 第二次做
func findRepeatNumber2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	i := 0
	nowNum := 0 // 下标i对应的当前数字
	n := -1     // 重复数字
	for i < len(nums) {
		nowNum = nums[i]
		if nowNum < 0 || nowNum > len(nums) {
			panic("数字必须在0~n-1范围内")
		}

		// 若当前数字和下标相等，说明它本身就在正确的位置
		if nowNum == i {
			i++
			continue
		}

		// 若目标下标的值已经在正确位置，说明当前值是重复数字
		if nums[nowNum] == nowNum {
			n = nowNum
			break
		}

		// 否则交换当前数字到正确的位置
		nums[i], nums[nowNum] = nums[nowNum], nums[i]
	}
	return n
}
