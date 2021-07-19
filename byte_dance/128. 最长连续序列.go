package main

// 128. 最长连续序列

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 请你设计并实现时间复杂度为O(n) 的算法解决此问题。

func main() {

}

// 1.使用插入排序，在排序过程中计算最长连续序列
// 2.依次遍历数组中每一个数字x，不断尝试寻找x+1,x+2...是否存在，这个寻找的过程，可以使用map来做到O(1)的时间复杂度
func longestConsecutive(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	res := 0

	hash := make(map[int]bool)
	for _, num := range nums {
		hash[num] = true
	}

	for num := range hash {
		if hash[num-1] {
			continue
		}

		i := 1
		for ; hash[num+i]; i++ {
		}

		res = max(res, i)
	}

	return res
}
