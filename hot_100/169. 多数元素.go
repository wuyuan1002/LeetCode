package main

// 169. 多数元素

// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

// func main() {

// }

// 同 offer 39
func majorityElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

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
