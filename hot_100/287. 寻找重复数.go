package main

// 287. 寻找重复数

// 给定一个包含n + 1 个整数的数组nums ，其数字都在 1 到 n之间（包括 1 和 n），可知至少存在一个重复的整数。
// 假设 nums 只有一个重复的整数 ，找出这个重复的数 。
//
// 你设计的解决方案必须不修改数组 nums 且只用常量级 O(1) 的额外空间。

func main() {

}

// 类似 offer 01, offer 53-2
// 若数组排序，则可以从后向前遍历数组或者使用二分法，来寻找第一个与下标不相等的数字
// 但是现在数组不排序，使用Hot100 141环形链表的解决思路，快慢指针
func findDuplicate(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	slow, fast := 0, 0 // 快慢指针
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			fast = 0
			for slow != fast {
				slow = nums[slow]
				fast = nums[fast]
			}
			return slow
		}
	}
}
