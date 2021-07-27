package main

import (
	"fmt"
)

// 剑指 Offer 39. 数组中出现次数超过一半的数字

// 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Printf("%d", majorityElement(nums))
}

// 求水王数
// 相当于一次删除两个不同的数字，有水王数时最后剩下的就是水王数，
// 但若没有水王数时，也可能剩下来，此时需要再遍历一次数组查看剩下来数字的出现次数，比如1,2,3,4,5会剩下5
func majorityElement(nums []int) int {
	veto, num := 0, nums[0]
	for _, n := range nums {
		if veto == 0 {
			num = n
		}
		if n == num {
			veto++
		} else {
			veto--
		}
	}
	return num
}

// 第二次做
func majorityElement1(nums []int) int {
	veto, num := 0, 0
	for _, n := range nums {
		if veto == 0 {
			num = n
		}
		if n == num {
			veto++
		} else {
			veto--
		}
	}
	return num
}
